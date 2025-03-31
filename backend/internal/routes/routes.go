package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/veetipihlava/shakki-peli/internal/handlers"
)

func SetupRoutes(e *echo.Echo) {

	e.POST("/game", handlers.HandleCreateGame)
	e.POST("/game/:game-id/join", handlers.JoinAsPlayer)

	e.GET("/ws/game", handlers.UpgradeConnection)

}
