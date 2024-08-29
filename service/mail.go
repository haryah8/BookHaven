package email

import (
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendTopupNotification(toEmail, userName, amount string) error {
	// Your SendGrid API key
	apiKey := "YOUR_SENDGRID_API_KEY"

	// Create a new SendGrid client
	client := sendgrid.NewSendClient(apiKey)

	// Create the email content
	from := mail.NewEmail("Your Company Name", "no-reply@yourdomain.com")
	subject := "Top-Up Successful"
	to := mail.NewEmail(userName, toEmail)
	plainTextContent := "Dear " + userName + ",\n\nYour top-up of " + amount + " has been successfully processed.\n\nThank you for using our service."
	htmlContent := "<p>Dear " + userName + ",</p><p>Your top-up of " + amount + " has been successfully processed.</p><p>Thank you for using our service.</p>"

	// Create a new mail message
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	// Send the email
	response, err := client.Send(message)
	if err != nil {
		log.Println("Error sending email:", err)
		return err
	}

	log.Println("Email sent successfully, status code:", response.StatusCode)
	return nil
}
