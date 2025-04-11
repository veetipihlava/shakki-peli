package websockets

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/veetipihlava/shakki-peli/internal/chess"
	"github.com/veetipihlava/shakki-peli/internal/connections"
	"github.com/veetipihlava/shakki-peli/internal/sessionstore"
)

type Message struct {
	Type    string      `json:"type"`    // eg "move"
	Content interface{} `json:"content"` // eg ValidMoveContent
}

type ErrorContent struct {
	RefType string `json:"ref_type"` // eg "move"
	Error   string `json:"error"`    // eg "Move is invalid"
}

type JoinGameContent struct {
	UserID int64 `json:"user_id"`
	Color  bool  `json:"color"`
}

type LeaveGameContent struct {
	UserID int64 `json:"user_id"`
}

type ValidMoveContent struct {
	Move         string `json:"move"`
	UserID       int64  `json:"user_id"`
	KingInCheck  bool   `json:"king_in_check"`
	Draw         bool   `json:"draw"`
	Checkmate    bool   `json:"checkmate"`
	KingConsumed bool   `json:"king_consumed"`
	WinnerColor  bool   `json:"winner_color"`
}

// handleJoinGame handles verifying that user and game are in memory before linking this websocket connection to a gameID
// The join message is broadcasted to all players in game
func HandleJoinGame(ss sessionstore.SessionStore, ws *websocket.Conn, request ChessMessage) error {
	gameID := request.GameID
	userID := request.UserID

	// Verify player and game exists in memory
	_, player, err := GetGameAndPlayerFromSessionStore(ss, gameID, userID)
	if err != nil {
		msg := NewErrorMessage("join", "Error verifying player")
		Respond(ws, msg)
		return err
	}

	// Link this websocket connection to game and player
	err = connections.ConnManager.AddPlayerConnection(gameID, player, ws)
	if err != nil {
		log.Printf("[WS] Error adding playerConn: %v", err)
		msg := NewErrorMessage("join", "Error joining game")
		Respond(ws, msg)
		return err
	}

	// Broadcast join message to all players
	content := JoinGameContent{
		UserID: userID,
		Color:  player.Color,
	}
	msg := NewMessage("join", content)
	Broadcast(gameID, msg)
	return nil
}

// handleMoveRequest processes a move request from a player
func HandleMove(ss sessionstore.SessionStore, ws *websocket.Conn, request ChessMessage) error {
	gameID := request.GameID
	notation := request.Content

	// Verify player and game exist in memory
	_, player, err := GetGameAndPlayerFromSessionStore(ss, request.GameID, request.UserID)
	if err != nil {
		msg := NewErrorMessage("move", "Error verifying player")
		Respond(ws, msg)
		return err
	}

	// Fetch the pieces from memory
	pieces, err := ss.ReadPieces(gameID)
	if err != nil {
		return err
	}

	// Chess validator to check if move is valid
	validationResult, updatePieces := chess.ValidateMove(pieces, notation, player.Color)
	if !validationResult.IsValidMove {
		log.Println(validationResult)
		msg := NewErrorMessage("move", "Move not valid")
		Respond(ws, msg)
		return nil
	}

	// Save the valid move to stack
	validMove := GetAsMove(gameID, validationResult.Move)
	err = ss.SaveMove(validMove)
	if err != nil {
		msg := NewErrorMessage("move", "Error occured saving move")
		Respond(ws, msg)
		return nil
	}

	// Create a ChessEntry from the valid move
	chessEntry := GetAsChessEntry(gameID, validMove, validationResult.GameOver, updatePieces)
	err = ss.PublishEntry(chessEntry)
	if err != nil {
		return nil
	}

	// Broadcast the valid move
	content := GetAsValidMoveContent(player.UserID, validationResult)
	msg := NewMessage("move", content)
	Broadcast(gameID, msg)
	return nil
}

// handleClosing processes a closing request from a player
func HandleLeave(ss sessionstore.SessionStore, ws *websocket.Conn) error {

	// Fetch the corresponding player and game from the conneciton
	player, err := connections.ConnManager.GetPlayerByConnection(ws)
	if err != nil {
		return err
	}

	// Remove connection
	err = connections.ConnManager.RemoveConnection(ws)
	if err != nil {
		return err
	}

	conns, err := connections.ConnManager.GetConnectionsInGame(player.GameID)
	if err != nil {
		return err
	}

	if len(conns) == 0 {
		connections.ConnManager.RemoveGame(player.GameID)
	}

	// Cleanup memory
	RemovePlayerFromSessionStore(ss, player.GameID, player.UserID)

	content := LeaveGameContent{UserID: player.UserID}
	msg := NewMessage("closing", content)
	Broadcast(player.GameID, msg)

	return nil
}
