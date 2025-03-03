package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/veetipihlava/shakki-peli/internal/database"
	"github.com/veetipihlava/shakki-peli/internal/games"
	"github.com/veetipihlava/shakki-peli/internal/middleware"
	"github.com/veetipihlava/shakki-peli/internal/utilities"
)

type CreateGamesResponse struct {
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

	response := CreateGamesResponse{
		GameID: gameID,
	}

	return c.JSON(http.StatusOK, response)
}
