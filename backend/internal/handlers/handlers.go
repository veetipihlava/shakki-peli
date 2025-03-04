package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/veetipihlava/shakki-peli/internal/chess"
	"github.com/veetipihlava/shakki-peli/internal/database"
	"github.com/veetipihlava/shakki-peli/internal/games"
	"github.com/veetipihlava/shakki-peli/internal/middleware"
	"github.com/veetipihlava/shakki-peli/internal/utilities"
)

type CreateGameResponse struct {
	GameID int64 `json:"game_id"`
}

// Creates a new game and redirects user.
func HandleCreateGame(c echo.Context) error {
	db := c.Get(middleware.DatabaseContextName).(*database.DatabaseService)
	gameID, err := utilities.CreateNewChessGame(db)
	if err != nil {
		return err
	}

	games.GameManager.CreateGame(gameID)
	log.Printf("Created new game: %d", gameID)

	response := CreateGameResponse{
		GameID: gameID,
	}

	return c.JSON(http.StatusOK, response)
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

type JoinGameResponse struct {
	PlayerID int64 `json:"player_id"`
}

// Creates a new player to the game.
func HandleJoinGame(c echo.Context) error {
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
		return err
	}

	log.Printf("%d joined game %d", playerID, gameID)

	response := JoinGameResponse{
		PlayerID: playerID,
	}

	return c.JSON(http.StatusOK, response)
}
