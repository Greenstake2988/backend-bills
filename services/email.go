package services

import (
	"fmt"
	"net/smtp"
	"strconv"

	"github.com/spf13/viper"
)

func SendVerificationCodeEmail(userEmail string, verificationCode string) error {
	// Set up the SMTP server configuration
	smtpServer := viper.GetString("SMTP.SERVER")
	smtpPort := viper.GetInt("SMTP.PORT")
	smtpUsername := viper.GetString("SMTP.USERNAME")
	smtpPassword := viper.GetString("SMTP.PASSWORD")

	// Set up the email message
	from := smtpUsername
	to := userEmail
	subject := "Register Billie Bills Account"

	body := fmt.Sprintf("This is your verification code: %s", verificationCode)

	message := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body)

	// Convert the integer to a string
	strPort := strconv.Itoa(smtpPort)
	// Connect to the SMTP server
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpServer)
	err := smtp.SendMail(smtpServer+":"+strPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
