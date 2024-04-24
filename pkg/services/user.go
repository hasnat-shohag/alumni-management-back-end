package services

import (
	"alumni-management-server/pkg/domain"
	"alumni-management-server/pkg/email"
	"fmt"
)

type userService struct {
	userRepo domain.IUserRepo
	authRepo domain.IAuthRepo
}

func UserServiceInstance(userRepo domain.IUserRepo, authRepo domain.IAuthRepo) domain.IUserService {
	return &userService{
		userRepo: userRepo,
		authRepo: authRepo,
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
	link := fmt.Sprintf("http://localhost:9030/reset-password?otp=%s", otp)
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
