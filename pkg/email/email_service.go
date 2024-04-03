package email

import (
	"alumni-management-server/pkg/config"
	"fmt"
	"net/smtp"
)

func SendEmail(to, subject string, msg string) error {
	smtpConfig := config.LocalSMTPConfig

	auth := smtp.PlainAuth("", smtpConfig.Username,
		smtpConfig.Password, smtpConfig.Host)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	email := []byte(fmt.Sprintf("Subject: %s\r\n %s\r\n\r\n %s", subject, headers, msg))
	//email := "Subject: " + subject + "\r\n" + headers + "\r\n\r\n" + msg

	err := smtp.SendMail(
		fmt.Sprintf("%s:%s", smtpConfig.Host, smtpConfig.Port),
		auth,
		smtpConfig.Username,
		[]string{to},
		email,
		//[]byte(email),
	)

	if err != nil {
		return err
	}

	return nil
}
