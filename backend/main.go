package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/veetipihlava/shakki-peli/internal/database"
	"github.com/veetipihlava/shakki-peli/internal/handlers"
	"github.com/veetipihlava/shakki-peli/internal/middleware"
)

func main() {
	db, connection, err := database.CreateTestLibSQLConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer connection.Close()

	e := echo.New()
	e.Use(middleware.WithContext(middleware.DatabaseContextName, db))
	e.Use(middleware.UseUser) // TODO should? be replaced with some authentication stuff

	e.POST("/game", handlers.HandleCreateGame)
	e.GET("/game/:game-id", handlers.UpgradeJoinGame)

	e.Logger.Fatal(e.Start(":8080"))
}
