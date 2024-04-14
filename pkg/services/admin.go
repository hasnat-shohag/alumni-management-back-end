package services

import "alumni-management-server/pkg/domain"

type adminService struct {
	adminRepo domain.IAdminRepo
}

// NewAdminService returns a new instance of the adminService struct.
func NewAdminService(adminRepo domain.IAdminRepo) domain.IAdminService {
	return &adminService{
		adminRepo: adminRepo,
	}
}

func (adminService *adminService) VerifyUser(studentId string) error {
	// pass the request to the repository layer
	if err := adminService.adminRepo.VerifyUser(studentId); err != nil {
		return err
	}

	return nil
}
