package handler

import (
	"BookHaven/models"
	"BookHaven/service"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// Jokes.
// @Summary Papa's Joke
// @Description Jokes to brighten your day :D.
// @Tags Welcome
// @Produce json
// @Success 200 {array} models.JokeDto
// @Failure 500 {object} models.ErrorResponse
// @Router /books [get]
func WelcomeJoke(db *sql.DB, logger *logrus.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		joke, err := service.FetchRandomJoke()
		if err != nil {
			return c.JSON(http.StatusOK, models.ErrorInternalServer)
		}

		response := c.JSON(http.StatusOK, models.JokeDto{
			Joke: joke,
		})

		// Respond with success
		return response
	}
}
