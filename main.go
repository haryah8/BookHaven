package main

import (
	"BookHaven/config"
	"BookHaven/database"
	"BookHaven/handler"
	"BookHaven/logger"
	"BookHaven/middlewares"
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	logger := logger.InitLogger()
	config.InitConfig(logger)
	database.InitDB(logger)
	db := database.GetDB()
	defer database.CloseDB()

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middlewares.PrintRequestResponse(logger))

	e.POST("/", handler.PingMe(db, logger))
	e.POST("/login", handler.Login(db, logger))
	e.POST("/register", handler.Register(db, logger))

	// Apply JWT middleware to the protected route
	e.GET("/protected", handler.ProtectedEndpoint, middlewares.JWT(logger))

	// user := e.Group("/v1")
	// user.User(middlewares.JWT())
	var port string
	if config.PORT == "" {
		port = fmt.Sprintf(":%s", config.PORT)
	} else {
		port = ":8085"
	}

	e.Start(port)

}
