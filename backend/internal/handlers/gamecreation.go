package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/veetipihlava/shakki-peli/internal/chess"
	"github.com/veetipihlava/shakki-peli/internal/connections"
	"github.com/veetipihlava/shakki-peli/internal/database"
	"github.com/veetipihlava/shakki-peli/internal/middleware"
	"github.com/veetipihlava/shakki-peli/internal/models"
	"github.com/veetipihlava/shakki-peli/internal/sessionstore"
)

type CreateGameResponse struct {
	GameID int64 `json:"game_id"`
}

// Creates a new game and pieces to PersistentStorage and saves them to SessionStore.
func HandleCreateGame(c echo.Context) error {
	db := c.Get(middleware.DatabaseContextName).(*database.DatabaseService)
	redis := c.Get(middleware.RedisContextName).(sessionstore.SessionStore)

	// Initialize a new chess game
	game, err := db.CreateGame()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	// Initialize chess pieces
	pieces, err := db.CreatePieces(game.ID, chess.GetInitialChessGamePieces(game.ID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	connections.ConnManager.AddGame(game.ID)

	redis.SaveGame(game)
	redis.SavePieces(pieces)

	response := CreateGameResponse{
		GameID: game.ID,
	}

	return c.JSON(http.StatusOK, response)
}

// Verifies that a game with provided id is ongoing.
func validateGame(gameID int64) error {
	err := connections.ConnManager.VerifyGame(gameID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "No such game found")
	}
	return nil
}

type JoinGameResponse struct {
	UserID int64 `json:"user_id"`
}

// Creates a new player to PersistentStorage and saves them to SessionStore
func JoinAsPlayer(c echo.Context) error {
	user := c.Get(middleware.UserContextName).(*models.User)
	db := c.Get(middleware.DatabaseContextName).(*database.DatabaseService)
	redis := c.Get(middleware.RedisContextName).(sessionstore.SessionStore)

	gameIDString := c.Param("game-id")
	gameID, err := strconv.ParseInt(gameIDString, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid game ID format")
	}

	err = validateGame(gameID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "No ongoing game with provided id")
	}

	player, err := db.CreatePlayer(user.ID, gameID)
	if err != nil {
		return err
	}

	redis.SavePlayer(player)

	response := JoinGameResponse{
		UserID: player.UserID,
	}

	return c.JSON(http.StatusOK, response)
}
