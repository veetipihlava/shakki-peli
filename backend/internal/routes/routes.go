package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/veetipihlava/shakki-peli/internal/handlers"
	"github.com/veetipihlava/shakki-peli/internal/websockets"
)

func SetupRoutes(e *echo.Echo) {

	e.POST("/game", handlers.HandleCreateGame)
	e.POST("/game/:game-id/join", handlers.JoinAsPlayer)

	e.GET("/games/:game-id", handlers.HandleGetFullGame)

	e.GET("/ws", websockets.UpgradeConnection)
}
