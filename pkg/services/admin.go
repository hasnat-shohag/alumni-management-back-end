package services

import (
	"alumni-management-server/pkg/domain"
	"alumni-management-server/pkg/email"
)

type adminService struct {
	adminRepo domain.IAdminRepo
}

// NewAdminService returns a new instance of the adminService struct.
func NewAdminService(adminRepo domain.IAdminRepo) domain.IAdminService {
	return &adminService{
		adminRepo: adminRepo,
	}
}

func (adminService *adminService) VerifyUser(studentId string, isValid bool) error {
	// send verification successful email to the user
	user, err := adminService.adminRepo.FindUserByStudentId(studentId)
	if err != nil {
		return err
	}

	if isValid == true {
		err = email.SendEmail(user.Email, email.UserVerificationSuccess, email.UserVerificationSuccessTemplate)
		if err != nil {
			return err
		}
	} else {
		err = email.SendEmail(user.Email, email.UserVerificationFailed, email.UserVerificationFailedTemplate)
		if err != nil {
			return err
		}
	}

	// pass the request to the repository layer
	if err := adminService.adminRepo.VerifyUser(studentId, isValid); err != nil {
		return err
	}

	return nil
}
