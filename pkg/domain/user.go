package domain

import (
	"alumni-management-server/pkg/models"
)

type IUserRepo interface {
	CreateOTP(user *models.UserDetail) (string, error)
	UpdateUser(user *models.UserDetail) error
	FindAllAlumni(offset, limit int) ([]models.UserDetail, error)
	FindUser(id string) (*models.UserDetail, error)
	CountAuthorizedUser() (int, error)
}

type IUserService interface {
	ForgetPassword(email string) error
	ResetPassword(user *models.UserDetail, password string) error
	GetUserFromEmailWithValidOtp(email, otp string) (*models.UserDetail, error)
	GetAllAlumni(page, limit int) ([]models.UserDetail, int, error)
	GetUser(id string) (*models.UserDetail, error)
}
