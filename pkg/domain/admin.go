package domain

import (
	"alumni-management-server/pkg/common/response"
	"alumni-management-server/pkg/models"
	"alumni-management-server/pkg/serializer"
)

type IAdminRepo interface {
	VerifyUser(studentId string, isValid bool) error
	FindUserByStudentId(studentId string) (models.UserDetail, error)
	DeleteUser(user *models.UserDetail) error
	AddExecutiveCommitteeMember(executiveCommitteeMember *models.ExecutiveCommittee) error
	DeleteExecutiveCommitteeMember(execCommitteeMember *models.ExecutiveCommittee) error
	UpdateExecutiveCommitteeMember(execCommitteeMember *models.ExecutiveCommittee) error
	GetExecutiveCommitteeInfo() ([]models.ExecutiveCommittee, error)
	FindBy(fieldName string, value any) (models.ExecutiveCommittee, error)
}

type IAdminService interface {
	VerifyUser(studentId string, isValid bool) error
	DeleteUser(studentId string) error
	AddExecutiveCommitteeMember(request *serializer.CreateCommitteeRequest) error
	DeleteExecutiveCommitteeMember(id string) error
	UpdateExecutiveCommitteeMember(id string, request *serializer.UpdateCommitteeRequest) error
	GetExecutiveCommitteeInfo() ([]response.ExecutiveCommitteeResponse, error)
}
