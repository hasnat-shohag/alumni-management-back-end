package domain

import "alumni-management-server/pkg/models"

type IAdminRepo interface {
	VerifyUser(studentId string, isValid bool) error
	FindUserByStudentId(studentId string) (models.UserDetail, error)
}

type IAdminService interface {
	VerifyUser(studentId string, isValid bool) error
}
