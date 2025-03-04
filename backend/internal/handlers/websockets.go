package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/veetipihlava/shakki-peli/internal/database"
	"github.com/veetipihlava/shakki-peli/internal/games"
	"github.com/veetipihlava/shakki-peli/internal/middleware"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ChessMessage struct {
	Type     string `json:"type"`
	GameID   int64  `json:"game_id"`
	PlayerID int64  `json:"player_id"`
	Content  string `json:"content"`
}

// UpgradeJoinGame handles WebSocket connection for game participation
func UpgradeJoinGame(c echo.Context) error {
	db := c.Get(middleware.DatabaseContextName).(*database.DatabaseService)
	gameID, err := validateGameID(c)
	if err != nil {
		return err
	}

	ws, err := Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("WebSocket read error: %v", err)
			games.GameManager.EndGame(gameID)
			return nil
		}

		var request ChessMessage
		err = json.Unmarshal(msg, &request)
		if err != nil {
			log.Printf("invalid message: %v", err)
			sendErrorMessage(ws, "invalid message")
			continue
		}

		switch request.Type {
		case "join":
			handleJoinRequest(gameID, ws, request)
		case "move":
			handleMoveRequest(db, gameID, request)
		case "resign":
			// TODO: Implement resignation logic
		case "draw_offer":
			// TODO: Implement draw offer logic
		case "draw_response":
			// TODO: Implement draw response logic
		default:
			log.Printf("Unknown request type: %s", request.Type)
			sendErrorMessage(ws, "Unknown request type")
		}
	}
}
