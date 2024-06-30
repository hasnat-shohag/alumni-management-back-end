package services

import (
	"alumni-management-server/pkg/domain"
	"alumni-management-server/pkg/email"
	"alumni-management-server/pkg/models"
	"alumni-management-server/pkg/types"
	"alumni-management-server/pkg/utils"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
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

	// Open the Certificate or Student id Card file
	file, err := registerRequest.CertificateOrIdCard.Open()
	if err != nil {
		return err
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	// Create a new file in the desired location
	dirPath := "./images/certificates_and_id_cards"
	imagePath := filepath.Join(dirPath, registerRequest.StudentId+"_"+registerRequest.CertificateOrIdCard.Filename)

	// Create the directory if it doesn't exist
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			return err
		}
	}

	dst, err := os.Create(imagePath)
	if err != nil {
		return err
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			return
		}
	}(dst)

	// Copy the uploaded file to the new file
	if _, err := io.Copy(dst, file); err != nil {
		return err
	}

	// create user
	user := &models.UserDetail{
		Name:                           registerRequest.Name,
		StudentId:                      registerRequest.StudentId,
		Email:                          registerRequest.Email,
		GraduationYear:                 registerRequest.GraduationYear,
		CertificateOrStudentIdCardPath: imagePath,
		Role:                           registerRequest.Role,
		PasswordHash:                   passwordHash,
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
	accessToken, err := utils.GetJwtForUser(user)
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
