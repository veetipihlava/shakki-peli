package websockets

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/veetipihlava/shakki-peli/internal/connections"
	"github.com/veetipihlava/shakki-peli/internal/models"
	"github.com/veetipihlava/shakki-peli/internal/sessionstore"
)

func NewMessage(msgType string, content interface{}) Message {
	return Message{
		Type:    msgType,
		Content: content,
	}
}

func NewErrorMessage(refType string, message string) Message {
	return Message{
		Type: "error",
		Content: ErrorContent{
			RefType: refType,
			Error:   message,
		},
	}
}

func SendErrorMessage(conn *websocket.Conn, message string) {
	errorResponse := map[string]string{"error": message}
	Respond(conn, errorResponse)
}

// Respond sends a message to the connection.
func Respond(conn *websocket.Conn, response interface{}) {
	responseJSON, err := json.Marshal(response)
	if err != nil {
		log.Printf("Failed to marshal response: %v", err)
		return
	}

	err = conn.WriteMessage(websocket.TextMessage, responseJSON)
	if err != nil {
		log.Printf("Failed to send response: %v", err)
	}
}

// Broadcast sends a message to all connections in the game.
func Broadcast(gameID int64, response interface{}) {
	conns, err := connections.ConnManager.GetConnectionsInGame(gameID)
	if err != nil {
		log.Printf("Error fetching all connections from game %v: %v", gameID, err)
	}

	for _, conn := range conns {
		Respond(conn, response)
	}
}

// Get Game and Player from Session Store.
func GetGameAndPlayerFromSessionStore(ss sessionstore.SessionStore, gameID int64, userID int64) (*models.Game, *models.Player, error) {
	game, err := ss.ReadGame(gameID)
	if err != nil {
		log.Printf("[SESSION STORE] Error retrieving game %v from session store: %v", gameID, err)
		return nil, nil, err
	}
	player, err := ss.ReadPlayer(userID, gameID)
	if err != nil {
		log.Printf("[SESSION STORE] Error retrieving player %v from session store: %v", userID, err)
		return nil, nil, err
	}

	return game, player, nil
}

// Remove Player from game in memory. If game is empty, remove game and pieces also.
func RemovePlayerFromSessionStore(ss sessionstore.SessionStore, gameID int64, userID int64) error {

	_, _, err := GetGameAndPlayerFromSessionStore(ss, gameID, userID)
	if err != nil {
		log.Printf("[SESSION STORE] Game or player not found in memory. Continuing with cleanup")
	}

	// Remove the player
	if err := ss.RemovePlayer(userID, gameID); err != nil {
		log.Printf("[SESSION STORE] Error removing player %v from session store: %v", userID, err)
	}

	// Check if there are still players
	players, err := ss.ReadPlayers(gameID)
	if err != nil {
		log.Printf("[SESSION STORE] Error reading players from game %d: %v", gameID, err)
		log.Printf("[SESSION STORE] Proceeding to remove the game still.")
	}

	if err != nil || len(players) == 0 {
		log.Printf("[WS] Removing game %v from session store.", gameID)
		if err := ss.RemoveGame(gameID); err != nil {
			log.Printf("[WS] Error removing game %v: %v", gameID, err)
			return err
		}
	}

	return nil
}

func GetAsValidMoveContent(userID int64, result models.ValidationResult) ValidMoveContent {
	return ValidMoveContent{
		Move:         result.Move,
		UserID:       userID,
		KingInCheck:  result.KingInCheck,
		Draw:         result.GameOver.Draw,
		Checkmate:    result.GameOver.Checkmate,
		KingConsumed: result.GameOver.KingConsumed,
		WinnerColor:  result.GameOver.WinnerColor,
	}
}

func GetAsMove(gameID int64, notation string) models.Move {
	return models.Move{
		ID:        0,
		GameID:    gameID,
		Notation:  notation,
		CreatedAt: time.Now(),
	}
}

func GetAsChessEntry(gameID int64, move models.Move, gameOver models.GameOver, affectedPieces []models.PieceUpdate) models.ChessEntry {
	return models.ChessEntry{
		GameID:         gameID,
		Move:           move,
		GameOver:       gameOver,
		AffectedPieces: affectedPieces,
	}
}
