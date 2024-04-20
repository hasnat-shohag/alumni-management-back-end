package domain

import "alumni-management-server/pkg/models"

type IAdminRepo interface {
	VerifyUser(studentId string) error
	FindUserByStudentId(studentId string) (models.UserDetail, error)
}

type IAdminService interface {
	VerifyUser(studentId string) error
}
