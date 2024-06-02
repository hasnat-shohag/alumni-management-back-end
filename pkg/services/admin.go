package services

import (
	"alumni-management-server/pkg/domain"
	"alumni-management-server/pkg/email"
	"alumni-management-server/pkg/models"
	"fmt"
)

type adminService struct {
	adminRepo domain.IAdminRepo
	authRepo  domain.IAuthRepo
}

// NewAdminService returns a new instance of the adminService struct.
func NewAdminService(adminRepo domain.IAdminRepo, authRepo domain.IAuthRepo) domain.IAdminService {
	return &adminService{
		adminRepo: adminRepo,
		authRepo:  authRepo,
	}
}

func (adminService *adminService) VerifyUser(studentId string, isValid bool) error {
	// check the user is already verified or admin

	// send verification successful email to the user
	user, err := adminService.adminRepo.FindUserByStudentId(studentId)
	if err != nil {
		return err
	}
	// if user is admin then return error coz admin has no access to verify another admin
	if user.Role != "user" {
		return fmt.Errorf("user not found")
	}

	if user.IsUserVerified == true {
		return fmt.Errorf("user already verified")
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

func (adminService *adminService) DeleteUser(studentId string) error {
	// Find Authorized User
	var user *models.UserDetail
	user, err := adminService.authRepo.FindAuthorizedUserByEmailOrStudentId(studentId)
	if err != nil {
		return err
	}

	if err := adminService.adminRepo.DeleteUser(user); err != nil {
		return err
	}
	return nil
}
