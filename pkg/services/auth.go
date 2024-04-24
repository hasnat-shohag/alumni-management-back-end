package services

import (
	"alumni-management-server/pkg/domain"
	"alumni-management-server/pkg/email"
	"alumni-management-server/pkg/models"
	"alumni-management-server/pkg/types"
	"alumni-management-server/pkg/utils"
	"time"
)

// authService defines the methods of the domain.IAuthService interface.
type authService struct {
	authRepo domain.IAuthRepo
}

// AuthServiceInstance returns a new instance of the authService struct.
func AuthServiceInstance(authRepo domain.IAuthRepo) domain.IAuthService {
	return &authService{
		authRepo: authRepo,
	}
}

// SignupUser creates a new user with the given user details.
func (service *authService) SignupUser(registerRequest *types.SignupRequest) error {
	// Check if the user already exists
	err := service.authRepo.DuplicateUserChecker(&registerRequest.StudentId, &registerRequest.Email)
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
	user.OtpExpiryTime = time.Now()
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

	if err := service.authRepo.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func (service *authService) Login(loginRequest *types.LoginRequest) (*types.LoginResponse, error) {
	// Check user is verified or not
	var identifier *string
	// if studentId or email is not provided it gets an error from the validation in the controller layer
	if loginRequest.Email != nil {
		identifier = loginRequest.Email
	} else {
		identifier = loginRequest.StudentId
	}

	user, err := service.authRepo.FindAuthorizedUserByEmailOrStudentId(identifier)
	if err != nil {
		return nil, err
	}

	// Check password
	if err := utils.CheckPassword(user.PasswordHash, loginRequest.Password); err != nil {
		return nil, err
	}

	// Create JWT token
	accessToken, err := utils.GetJwtForUser(user.StudentId)
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		Name:           user.Name,
		Email:          user.Email,
		StudentId:      user.StudentId,
		IsUserVerified: user.IsUserVerified,
		IsActive:       true,
		Role:           user.Role,
		AccessToken:    accessToken,
	}, nil
}
