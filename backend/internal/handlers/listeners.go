package handlers

import (
	"encoding/json"
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

// sendMessageToAllPlayers sends a response to all players in a game
func sendMessageToAllPlayers(gameID int64, response interface{}) {
	players := games.GameManager.GetPlayers(gameID)
	for _, player := range players {
		sendResponse(player.Connection, response)
	}
}

// Send an error message to the client
func sendErrorMessage(conn *websocket.Conn, message string) {
	errorResponse := map[string]string{"error": message}
	sendResponse(conn, errorResponse)
}

type JoinResponse struct {
	Name string `json:"content"`
}

// handleJoinRequest processes a join request from a player
func handleJoinRequest(gameID int64, ws *websocket.Conn, request ChessMessage) {
	player := games.Player{
		Name:       request.Content,
		ID:         request.PlayerID,
		Connection: ws,
	}

	games.GameManager.AddPlayerToGame(gameID, player)

	response := JoinResponse{
		Name: player.Name,
	}

	sendMessageToAllPlayers(gameID, response)
}

type ValidationResponse struct {
	Move             string                 `json:"move"`
	ValidationResult chess.ValidationResult `json:"validation_result"`
}

// handleMoveRequest processes a move request from a player
func handleMoveRequest(db *database.DatabaseService, gameID int64, request ChessMessage) {
	validationResult, err := utilities.ProcessChessMove(db, request.PlayerID, gameID, request.Content)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	response := ValidationResponse{
		Move:             request.Content,
		ValidationResult: validationResult,
	}

	sendMessageToAllPlayers(gameID, response)
}
