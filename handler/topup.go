package handler

import (
	"BookHaven/config"
	"BookHaven/models"
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func TopUp(db *sql.DB, logger *logrus.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request models.TopUpRequestDto
		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
		}

		invoiceURL := "https://api.xendit.co/v2/invoices"
		xenditAPIKey := config.XENDIT_SECRET_KEY

		// Generate a new UUID and set up the invoice request body
		invoiceReqBody := models.TopUpRequestXenditDto{
			ExternalID: uuid.New().String(),
			Amount:     request.Amount,
		}
		logger.Info(fmt.Sprintf("Invoice Request: %v", invoiceReqBody)) // -------------------------------------- FOR DEBUGGING --------------------------------

		// Marshal request body into JSON
		reqBody, err := json.Marshal(invoiceReqBody)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create request body"})
		}

		client := &http.Client{}
		httpReq, err := http.NewRequest("POST", invoiceURL, bytes.NewBuffer(reqBody))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create request"})
		}

		httpReq.Header.Set("Authorization", "Basic "+basicAuth(xenditAPIKey))
		httpReq.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(httpReq)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to send request"})
		}
		defer resp.Body.Close()

		// Log the response status
		logger.Infof("Response status from Xendit: %s", resp.Status)

		// Read and Unmarshal the response body
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to read response"})
		}

		var response map[string]interface{}
		if err := json.Unmarshal(respBody, &response); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to parse response"})
		}
		logger.Info(fmt.Sprintf("Response From Xendit: %v", response)) // -------------------------------------- FOR DEBUGGING --------------------------------

		// Check for errors in the response
		if errorCode, ok := response["error_code"].(string); ok {
			errorMessage := "An error occurred"
			if msg, ok := response["message"].(string); ok {
				errorMessage = msg
			}

			// Return error response
			return c.JSON(http.StatusBadRequest, models.TopUpErrorResponseXenditDto{
				ErrorCode:    errorCode,
				ErrorMessage: errorMessage,
			})
		}

		// Extract only the required fields if no errors
		topUpResponse := models.TopUpResponseXenditDto{
			Status:     response["status"].(string),
			Amount:     fmt.Sprintf("%.0f", math.Round(response["amount"].(float64))), // Convert to string
			InvoiceURL: response["invoice_url"].(string),
			ExpiryDate: response["expiry_date"].(string),
		}
		user := c.Get("user")

		claims, ok := user.(*models.Claims) // Using type assertion
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid user claims"})
		}
		fmt.Println(claims.UserId)
		fmt.Println(claims.Email)
		// Insert user into database
		_, err = db.Exec(`INSERT INTO transactions (user_id, amount, transaction_id, description) VALUES (?, ?, ?, ?)`,
			claims.UserId, invoiceReqBody.Amount, invoiceReqBody.ExternalID, "TopUp Balance")
		if err != nil {
			logger.Error("Error inserting payment data into database: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"status":  "Error",
				"message": "Failed to insert payment data",
			})
		}

		return c.JSON(http.StatusOK, topUpResponse)
	}
}

func basicAuth(apiKey string) string {
	return base64.StdEncoding.EncodeToString([]byte(apiKey + ":"))
}
