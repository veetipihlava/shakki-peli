package handlers

import (
	"errors"
	"log"

	"github.com/gorilla/websocket"
	"github.com/veetipihlava/shakki-peli/internal/connections"
	"github.com/veetipihlava/shakki-peli/internal/models"
	"github.com/veetipihlava/shakki-peli/internal/sessionstore"
	"github.com/veetipihlava/shakki-peli/internal/utilities"
)

type Message struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

// User joins the WebSocket. Game, pieces and player are already in SessionStore.
func joinWebSocket(redis sessionstore.SessionStore, ws *websocket.Conn, request ChessMessage) error {
	_, err := redis.ReadGame(request.GameID)
	if err != nil {
		return errors.New("game not in sessionstore")
	}
	player, err := redis.ReadPlayer(request.PlayerID, request.GameID)
	if err != nil {
		return errors.New("player not in sessionstore")
	}

	playerConn := connections.PlayerConn{Player: player, Conn: ws}

	err = connections.ConnManager.TryAddPlayerToGame(request.GameID, playerConn)
	if err != nil {
		log.Printf("Could not add player: %v", err)
		return err
	}

	players, err := connections.ConnManager.GetGameConnections(request.GameID)

	if err != nil {
		log.Printf("Could not read players: %v", err)
		return err
	}

	message := Message{Type: "join", Content: request.Content}

	utilities.SendMessageToAllPlayers(players, request.GameID, message)

	return nil
}

// handleClosing processes a closing request from a player
func handleClosing(ws *websocket.Conn) error {
	gameID, playerConn, err := connections.ConnManager.GetPlayerFromConn(ws)

	err = connections.ConnManager.RemovePlayer(gameID, playerConn)
	if err != nil {
		log.Printf("Could not delete player: %v", err)
	}

	players, err := connections.ConnManager.GetGameConnections(gameID)
	if err != nil {
		log.Printf("Could not read players: %v", err)
	}

	message := Message{
		Type:    "closing",
		Content: "en tiedä vittu",
	}
	utilities.SendMessageToAllPlayers(players, gameID, message)

	return nil
}

// handleMoveRequest processes a move request from a player
func handleMoveRequest(request ChessMessage) error {

	valid := models.ValidationResult{
		Move:        request.Content,
		IsValidMove: true,
		KingInCheck: false,
		GameOver: models.GameOver{
			Draw:         false,
			Checkmate:    false,
			KingConsumed: false,
			WinnerColor:  true,
		},
	}

	players, err := connections.ConnManager.GetGameConnections(request.GameID)
	if err != nil {
		log.Printf("Could not read players: %v", err)
	}

	message := Message{
		Type:    "move",
		Content: valid.Move,
	}

	utilities.SendMessageToAllPlayers(players, request.GameID, message)
	return nil
}
