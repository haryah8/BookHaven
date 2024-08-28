package middlewares

import (
	"bytes"
	"io"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// PrintRequestResponse logs incoming requests and outgoing responses
func PrintRequestResponse(logger *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Log request details
			start := time.Now()
			req := c.Request()

			body, err := io.ReadAll(req.Body)
			if err != nil {
				logger.Errorf("Error reading request body: %v", err)
				return err
			}
			req.Body = io.NopCloser(bytes.NewBuffer(body))

			logger.Infof("Request: %s %s | Body: %s", req.Method, req.RequestURI, body)

			// Handle the request
			err = next(c)

			// Log response details
			resp := c.Response()

			duration := time.Since(start)

			logger.Infof("Response Status: %d | Duration: %v", resp.Status, duration)

			return err
		}
	}
}
