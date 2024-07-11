package domain

import (
	"alumni-management-server/pkg/models"
	"alumni-management-server/pkg/serializer"
)

type IAuthRepo interface {
	DuplicateUserChecker(StudentId *string, Email *string) error
	CreateUser(user *models.UserDetail) error
	FindAuthorizedUserByEmailOrStudentId(interface{}) (*models.UserDetail, error)
}

type IAuthService interface {
	SignupUser(registerRequest *serializer.SignupRequest) error
	Login(loginRequest *serializer.LoginRequest) (*serializer.LoginResponse, error)
}
