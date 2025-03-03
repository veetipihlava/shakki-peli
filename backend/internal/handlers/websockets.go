package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/veetipihlava/shakki-peli/internal/chess"
	"github.com/veetipihlava/shakki-peli/internal/database"
	"github.com/veetipihlava/shakki-peli/internal/games"
	"github.com/veetipihlava/shakki-peli/internal/middleware"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Validate the game ID from the request
func validateGameID(c echo.Context) (int64, error) {
	gameIDString := c.Param("game-id")
	gameID, err := strconv.ParseInt(gameIDString, 10, 64)
	if err != nil {
		return 0, echo.NewHTTPError(http.StatusBadRequest, "Invalid game ID format")
	}

	if !games.GameManager.GameExists(gameID) {
		return 0, echo.NewHTTPError(http.StatusNotFound, "No such game has been started")
	}

	if games.GameManager.GameIsFull(gameID) {
		return 0, echo.NewHTTPError(http.StatusUnauthorized, "The game is already full")
	}

	return gameID, nil
}

// UpgradeJoinGame handles WebSocket connection for game participation
func UpgradeJoinGame(c echo.Context) error {
	ws, err := Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	gameID, err := validateGameID(c)
	if err != nil {
		return err
	}

	userID := c.Get(middleware.UserIDContextName).(int64)
	db := c.Get(middleware.DatabaseContextName).(*database.DatabaseService)

	// TODO: Implement proper color assignment logic
	color := chess.White

	playerID, err := db.CreatePlayer(userID, gameID, color)
	if err != nil {
		log.Printf("Failed to create player: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to register player")
	}

	player := games.Player{
		Name:       "TODO", // TODO replace with actual user name
		ID:         playerID,
		Connection: ws,
	}

	games.GameManager.AddPlayerToGame(gameID, player)
	for {
		_, msg, err := player.Connection.ReadMessage()
		if err != nil {
			log.Printf("WebSocket read error: %v", err)
			games.GameManager.EndGame(gameID)
			return nil
		}

		var request ChessRequest
		err = json.Unmarshal(msg, &request)
		if err != nil {
			log.Printf("Failed to parse message: %v", err)
			sendErrorMessage(player.Connection, "Invalid message format")
			continue
		}

		switch request.Type {
		case "move":
			handleMoveRequest(request.Content, player, gameID, db)
		case "resign":
			// TODO: Implement resignation logic
		case "draw_offer":
			// TODO: Implement draw offer logic
		case "draw_response":
			// TODO: Implement draw response logic
		default:
			log.Printf("Unknown request type: %s", request.Type)
			sendErrorMessage(player.Connection, "Unknown request type")
		}
	}
}
