package domain

import (
	"alumni-management-server/pkg/models"
	"alumni-management-server/pkg/types"
)

type IAdminRepo interface {
	VerifyUser(studentId string, isValid bool) error
	FindUserByStudentId(studentId string) (models.UserDetail, error)
	DeleteUser(user *models.UserDetail) error
	AddExecutiveCommitteeMember(executiveCommitteeMember *models.ExecutiveCommittee) error
}

type IAdminService interface {
	VerifyUser(studentId string, isValid bool) error
	DeleteUser(studentId string) error
	AddExecutiveCommitteeMember(request *types.CreateCommitteeRequest) error
}
