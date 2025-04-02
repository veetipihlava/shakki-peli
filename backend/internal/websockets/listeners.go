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
}

type LeaveGameContent struct {
	UserID int64 `json:"user_id"`
}

type ValidMoveContent struct {
	Move         string `json:"move"`
	KingInCheck  string `json:"king_in_check"`
	GameOver     bool   `json:"game_over"`
	Draw         bool   `json:"draw"`
	Checkmate    bool   `json:"checkmate"`
	KingConsumed bool   `json:"king_consumed"`
	WinnerColor  bool   `json:"winner_color"`
}

// handleJoinGame handles verifying that user and game are in memory before linking this websocket connection to a gameID
// The join message is broadcasted to all players in game
func HandleJoinGame(ss sessionstore.SessionStore, ws *websocket.Conn, request ChessMessage) error {
	gameID := request.GameID
	userID := request.PlayerID

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
	content := JoinGameContent{UserID: userID}
	msg := NewMessage("join", content)
	Broadcast(gameID, msg)
	return nil
}

// handleMoveRequest processes a move request from a player
func HandleMove(ss sessionstore.SessionStore, ws *websocket.Conn, request ChessMessage) error {
	gameID := request.GameID
	//userID := request.PlayerID
	move := request.Content

	// Verify player and game exist in memory
	_, player, err := GetGameAndPlayerFromSessionStore(ss, request.GameID, request.PlayerID)
	if err != nil {
		return err
	}

	// Fetch the pieces from memory
	pieces, err := ss.ReadPieces(gameID)
	if err != nil {
		return err
	}

	// Call the Chess validator
	validationResult, _ := chess.ValidateMove(pieces, move, player.Color)

	if !validationResult.IsValidMove {
		msg := NewErrorMessage("move", "Move not valid")
		Respond(ws, msg)
		return nil
	}

	content := ValidMoveContent{}
	msg := NewMessage("move", content)
	Broadcast(gameID, msg)
	return nil
}

// handleClosing processes a closing request from a player
func HandleLeave(ss sessionstore.SessionStore, ws *websocket.Conn) error {

	// Fetch the corresponding player and game from the conneciton
	player, err := connections.ConnManager.GetPlayerByConnection(ws)

	// Remove connection
	err = connections.ConnManager.RemoveConnection(ws)
	if err != nil {
		log.Printf("Could not delete player: %v", err)
	}

	// Cleanup memory
	RemovePlayerFromSessionStore(ss, player.GameID, player.UserID)

	content := LeaveGameContent{UserID: player.UserID}
	msg := NewMessage("closing", content)
	Broadcast(player.GameID, msg)
	return nil
}
