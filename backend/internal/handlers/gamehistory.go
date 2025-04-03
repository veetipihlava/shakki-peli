package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/veetipihlava/shakki-peli/internal/database"
	"github.com/veetipihlava/shakki-peli/internal/middleware"
	"github.com/veetipihlava/shakki-peli/internal/models"
)

type FullGameResponse struct {
	Game    *models.Game    `json:"game"`
	Players []models.Player `json:"players"`
	Pieces  []models.Piece  `json:"pieces"`
	Moves   []models.Move   `json:"moves"`
}

func HandleGetFullGame(c echo.Context) error {
	db := c.Get(middleware.DatabaseContextName).(*database.DatabaseService)

	gameIDStr := c.Param("game-id")
	gameID, err := strconv.ParseInt(gameIDStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid game ID format")
	}

	game, players, pieces, moves, err := db.GetFullGameState(gameID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch game state")
	}
	if game == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Game not found")
	}

	response := FullGameResponse{
		Game:    game,
		Players: players,
		Pieces:  pieces,
		Moves:   moves,
	}

	return c.JSON(http.StatusOK, response)
}
