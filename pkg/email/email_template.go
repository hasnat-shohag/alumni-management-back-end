package email

import (
	"bytes"
	"html/template"
)

type AdminNotificationData struct {
	UserName  string
	StudentId string
}

const (
	UserVerificationSubject    = "Account Verification"
	AdminNotificationSubject   = "New User Registration"
	UserVerificationSuccess    = "User Verification Successful"
	UserVerificationsUnsuccess = "User Verification Unsuccessful"
	UserVerificationTemplate   = `
	Hello,
	
	Thank you for creating an account. Our admin will verify your information shortly. Please wait for the verification process to complete.
	
	Best regards,
	ICE alumni management Team
	`

	AdminNotificationTemplate = `
	<!DOCTYPE html>	
	<html>
	<head>
	  <style>
		body {
		  font-family: Arial, sans-serif;
		}
		p {
		  color: #333;
		}
		.highlight {
		  color: #ff0000;
		}
	  </style>
	<title>Verification Email</title>
	</head>
	<body>
	  <p>Hello Admin,</p>
	  <p>A new user, <span class="highlight">{{.UserName}}</span>, with Student Id <span class="highlight">{{.StudentId}}</span> has created an account and needs verification. Please review his/her information and verify the account.</p>
	  <p>Best regards,</p>
	  <p>ICE alumni management Team</p>
	</body>
	</html>
	`
	UserVerificationSuccessTemplate = `
		Dear User,
		Your account has been verified successfully. You can now login to your account.
	`
	UserVerficationUnsuccessTemplate = `
		Dear User,
		Your account verification was unsuccessful. Please create account with valid information.	
	`
)

func CreateAdminNotificationEmail(userName string, studentId string) (string, error) {
	data := AdminNotificationData{
		UserName:  userName,
		StudentId: studentId,
	}

	tmpl, err := template.New("AdminNotification").Parse(AdminNotificationTemplate)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, data); err != nil {
		return "", err
	}

	return tpl.String(), nil
}
