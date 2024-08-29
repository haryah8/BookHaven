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

	e.GET("/", handler.PingMe(db, logger))
	e.POST("/login", handler.Login(db, logger))
	e.POST("/register", handler.Register(db, logger))
	e.POST("/topup/callback", handler.TopUpCallback(db, logger))
	e.GET("/books", handler.GetAllBooks(db, logger))
	// Apply JWT middleware to the protected route
	e.GET("/protected", handler.ProtectedEndpoint, middlewares.JWT(logger))

	user := e.Group("/user")
	user.Use(middlewares.JWT(logger))

	user.GET("/balance", handler.GetUserBalance(db, logger))
	user.POST("/topup", handler.TopUp(db, logger))
	user.POST("/book/borrow", handler.BorrowBook(db, logger))
	user.POST("/book/return", handler.ReturnBook(db, logger))
	user.GET("/book/check", handler.GetBorrowedBooks(db, logger))

	port := fmt.Sprintf(":%s", config.PORT)
	if port == "" || port == ":" {
		port = ":8085"
	}
	// Start server
	e.Start(port)

}
