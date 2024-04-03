package services

import (
	"alumni-management-server/pkg/domain"
	"alumni-management-server/pkg/email"
	"alumni-management-server/pkg/models"
	"alumni-management-server/pkg/types"
	"alumni-management-server/pkg/utils"
)

// for email service

// authService defines the methods of the domain.IAuthService interface.
type authService struct {
	userRepo domain.IUserRepo
}

// AuthServiceInstance returns a new instance of the authService struct.
func AuthServiceInstance(userRepo domain.IUserRepo) domain.IAuthService {
	return &authService{
		userRepo: userRepo,
	}
}

// SignupUser creates a new user with the given user details.
func (service *authService) SignupUser(registerRequest *types.SignupRequest) error {
	// Check if the user already exists
	err := service.userRepo.DuplicateUserChecker(&registerRequest.StudentId, &registerRequest.Email)
	if err != nil {
		return err
	}

	// get hashed password
	passwordHash, err := utils.GetPasswordHash(registerRequest.Password)
	if err != nil {
		return err
	}

	// create user
	user := &models.UserDetail{
		StudentId:    registerRequest.StudentId,
		PasswordHash: passwordHash,
		Name:         registerRequest.Name,
		Email:        registerRequest.Email,
	}

	user.SetVerificationProperties()
	//? implement verification later

	//Send verification email to user
	err = email.SendEmail(user.Email, email.UserVerificationSubject, email.UserVerificationTemplate)
	if err != nil {
		return err
	}

	//////Notify admin
	emailBody, err := email.CreateAdminNotificationEmail(user.Name, user.StudentId)
	if err != nil {
		return err
	}

	adminEmail := "hasnat.ru.ice19@gmail.com"
	err = email.SendEmail(adminEmail, email.AdminNotificationSubject, emailBody)
	if err != nil {
		return err
	}

	if err := service.userRepo.CreateUser(user); err != nil {
		return err
	}

	return nil

}
