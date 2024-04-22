package domain

import (
	"alumni-management-server/pkg/models"
	"alumni-management-server/pkg/types"
)

type IUserRepo interface {
	DuplicateUserChecker(StudentId *string, Email *string) error
	CreateUser(user *models.UserDetail) error
	FindAuthorizedUserByEmailOrStudentId(interface{}) (*models.UserDetail, error)
}

type IAuthService interface {
	SignupUser(registerRequest *types.SignupRequest) error
	Login(loginRequest *types.LoginRequest) (*types.LoginResponse, error)
}
