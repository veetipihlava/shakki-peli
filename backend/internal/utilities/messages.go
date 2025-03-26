package utilities

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
	"github.com/veetipihlava/shakki-peli/internal/games"
)

// Send a response to the client
func SendResponse(conn *websocket.Conn, response interface{}) {
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

// Send an error message to the client
func SendErrorMessage(conn *websocket.Conn, message string) {
	errorResponse := map[string]string{"error": message}
	SendResponse(conn, errorResponse)
}

// sendMessageToAllPlayers sends a response to all players in a game
func SendMessageToAllPlayers(players []games.Player, gameID int64, response interface{}) {
	for _, player := range players {
		SendResponse(player.Connection, response)
	}
}
