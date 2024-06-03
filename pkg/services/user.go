package services

import (
	"alumni-management-server/pkg/domain"
	"alumni-management-server/pkg/email"
	"alumni-management-server/pkg/models"
	"alumni-management-server/pkg/types"
	"alumni-management-server/pkg/utils"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type userService struct {
	userRepo  domain.IUserRepo
	authRepo  domain.IAuthRepo
	adminRepo domain.IAdminRepo
}

func UserServiceInstance(userRepo domain.IUserRepo, authRepo domain.IAuthRepo, adminRepo domain.IAdminRepo) domain.IUserService {
	return &userService{
		userRepo:  userRepo,
		authRepo:  authRepo,
		adminRepo: adminRepo,
	}
}

func (userService *userService) ForgetPassword(Email string) error {
	// Check if user exists
	user, err := userService.authRepo.FindAuthorizedUserByEmailOrStudentId(Email)
	if err != nil {
		return fmt.Errorf("user with email %s does not exist", Email)
	}

	// Create OTP
	otp, err := userService.userRepo.CreateOTP(user)
	if err != nil {
		return err
	}

	// Send OTP to user
	link := fmt.Sprintf("http://localhost:9030/reset-password?otp=%s&email=%s", otp, Email)
	emailBody, err := email.CreateForgotPasswordEmail(link)
	if err != nil {
		return err
	}

	err = email.SendEmail(user.Email, email.PasswordResetSubject, emailBody)
	if err != nil {
		return err
	}
	return nil
}

func (userService *userService) GetUserFromEmailWithValidOtp(email, otp string) (*models.UserDetail, error) {
	user, err := userService.authRepo.FindAuthorizedUserByEmailOrStudentId(email)
	if err != nil {
		return nil, err
	}
	// check user otp is valid or not
	if err := utils.CheckPassword(user.OTP, otp); err != nil {
		return nil, fmt.Errorf("invalid otp")
	}

	// check otp expiry time
	if user.OtpExpiryTime.Before(user.OtpExpiryTime) {
		return nil, fmt.Errorf("otp expired, try again")
	}

	return user, nil
}

func (userService *userService) ResetPassword(user *models.UserDetail, password string) error {
	// Hash the password
	hashedPassword, err := utils.GetPasswordHash(password)
	if err != nil {
		return err
	}

	// Update the user password
	user.PasswordHash = hashedPassword
	err = userService.userRepo.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (userService *userService) GetAllAlumni(page, limit int) ([]models.UserDetail, int, error) {
	offset := (page - 1) * limit
	var alumni []models.UserDetail
	alumni, err := userService.userRepo.FindAllAlumni(offset, limit)

	if err != nil {
		return nil, 0, err
	}

	totalRecords, err := userService.userRepo.CountAuthorizedUser()
	if err != nil {
		return nil, 0, err
	}

	return alumni, totalRecords, nil
}

func (userService *userService) GetUser(id string) (*models.UserDetail, error) {
	user, err := userService.userRepo.FindUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (userService *userService) DeleteMe(studentId, studentIdFromToken string) error {
	if studentId != studentIdFromToken {
		return fmt.Errorf("you have no access to delete others account")
	}

	user, err := userService.userRepo.FindUser(studentId)
	if err != nil {
		return nil
	}

	err = userService.adminRepo.DeleteUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (userService *userService) UpdateMe(studentId string, request types.UpdateUserRequest) error {
	user, err := userService.userRepo.FindUser(studentId)
	if err != nil {
		return err
	}

	// If the user is not found, return an error
	if user == nil {
		return errors.New("user not found")
	}

	// Open the image file
	file, err := request.Image.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new file in the desired location
	dirPath := "./images"
	imagePath := filepath.Join(dirPath, request.Image.Filename)

	// Create the directory if it doesn't exist
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, 0755)
	}

	dst, err := os.Create(imagePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy the uploaded file to the new file
	if _, err := io.Copy(dst, file); err != nil {
		return err
	}

	// Update the user's image with the image from the request
	user.ImagePath = imagePath

	// Save the updated user back to the database
	err = userService.userRepo.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}
