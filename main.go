package main

import (
	"BookHaven/config"
	"BookHaven/database"
	"BookHaven/handler"
	"BookHaven/logger"
	"BookHaven/middlewares"
	"fmt"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "BookHaven/docs"

	"github.com/labstack/echo/v4/middleware"
)

// @title BookHaven API
// @version 1.0
// @description BookHaven Library - Travel around the world with books.
// @contact.name Harya Kumuda
// @contact.email hkkoostanto@students.hacktiv8.ac.id
// @BasePath /
// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
func main() {
	logger := logger.InitLogger()
	config.InitConfig(logger)
	database.InitDB(logger)
	db := database.GetDB()
	defer database.CloseDB()

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middlewares.PrintRequestResponse(logger))

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/", handler.WelcomeJoke(db, logger))
	e.GET("/ping", handler.PingMe(db, logger))
	e.POST("/login", handler.Login(db, logger))
	e.POST("/register", handler.Register(db, logger))
	e.POST("/topup/callback", handler.TopUpCallback(db, logger))
	e.GET("/books", handler.GetAllBooks(db, logger))
	// Apply JWT middleware to the protected route
	e.GET("/pingme", handler.ProtectedEndpoint, middlewares.JWT(logger))

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
