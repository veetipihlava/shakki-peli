package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/veetipihlava/shakki-peli/internal/database"
	"github.com/veetipihlava/shakki-peli/internal/middleware"
	"github.com/veetipihlava/shakki-peli/internal/routes"
	"github.com/veetipihlava/shakki-peli/internal/sessionstore/redis"
)

func main() {
	// Initialize the database
	db, connection, err := database.CreateTestLibSQLConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer connection.Close()

	// Initialize Session Storage
	redis, err := redis.InitializeRedis()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer redis.Close()
	e := echo.New()

	e.Use(middleware.WithContext(middleware.DatabaseContextName, db))
	e.Use(middleware.WithContext(middleware.RedisContextName, redis))
	e.Use(middleware.UseUser) // TODO should? be replaced with some authentication stuff

	// Create all endpoints
	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
