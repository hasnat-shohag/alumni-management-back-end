package domain

import "alumni-management-server/pkg/models"

type IUserRepo interface {
	CreateOTP(user *models.UserDetail) (string, error)
	UpdateUser(user *models.UserDetail) error
}

type IUserService interface {
	ForgetPassword(email string) error
	ResetPassword(user *models.UserDetail, password string) error
	GetUserFromEmailWithValidOtp(email, otp string) (*models.UserDetail, error)
}
