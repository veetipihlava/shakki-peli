package handlers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/veetipihlava/shakki-peli/internal/chess"
	"github.com/veetipihlava/shakki-peli/internal/database"
	"github.com/veetipihlava/shakki-peli/internal/games"
	"github.com/veetipihlava/shakki-peli/internal/utilities"
)

// Send a response to the client
func sendResponse(conn *websocket.Conn, response interface{}) {
	responseJSON, err := json.Marshal(response)
	if err != nil {
		log.Printf("Failed to marshal response: %v", err)
		return
	}

	if err := conn.WriteMessage(websocket.TextMessage, responseJSON); err != nil {
		log.Printf("Failed to send response: %v", err)
	}
}

// Send an error message to the client
func sendErrorMessage(conn *websocket.Conn, message string) {
	errorResponse := map[string]string{"error": message}
	sendResponse(conn, errorResponse)
}

type ChessRequest struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

type ValidationResponse struct {
	ValidationResult chess.ValidationResult `json:"validation_result"`
}

// Handle a move request
func handleMoveRequest(move string, player games.Player, gameID int64, db *database.DatabaseService) {
	fmt.Printf("\nprocessing player: %d\n", player.ID)
	validationResult, err := utilities.ProcessChessMove(db, player.ID, gameID, move)
	if err != nil {
		log.Printf("error: %v", err)
		sendErrorMessage(player.Connection, "Invalid move")
		return
	}

	response := ValidationResponse{
		ValidationResult: validationResult,
	}

	players := games.GameManager.GetPlayers(gameID)
	for _, player := range players {
		sendResponse(player.Connection, response)
	}
}
