package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/veetipihlava/shakki-peli/internal/middleware"
	"github.com/veetipihlava/shakki-peli/internal/sessionstore"
	"github.com/veetipihlava/shakki-peli/internal/utilities"
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
func UpgradeConnection(c echo.Context) error {
	redis := c.Get(middleware.RedisContextName).(sessionstore.SessionStore)

	ws, err := Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				handleClosing(ws)
			} else {
				log.Printf("WebSocket error: %v", err)
			}
			return nil
		}

		var request ChessMessage
		err = json.Unmarshal(msg, &request)
		if err != nil {
			log.Printf("invalid message: %v", err)
			utilities.SendErrorMessage(ws, "invalid message")
			continue
		}

		switch request.Type {
		case "join":
			err := joinWebSocket(redis, ws, request)
			if err != nil {
				utilities.SendErrorMessage(ws, err.Error())
				log.Printf("error joining game %d: %v", request.GameID, err)
			}
		case "move":
			err := handleMoveRequest(request)
			if err != nil {
				log.Printf("failed to handle move: %v", err)
			}
		case "leave":
			// TODO
		case "resign":
			// TODO: Implement resignation logic
		case "draw_offer":
			// TODO: Implement draw offer logic
		case "draw_response":
			// TODO: Implement draw response logic
		default:
			log.Printf("Unknown request type: %s", request.Type)
			utilities.SendErrorMessage(ws, "Unknown request type")
		}
	}
}
