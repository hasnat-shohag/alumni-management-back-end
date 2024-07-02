package services

import (
	"alumni-management-server/pkg/domain"
	"alumni-management-server/pkg/email"
	"alumni-management-server/pkg/models"
	"alumni-management-server/pkg/types"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type adminService struct {
	requestHashes map[string]bool
	adminRepo     domain.IAdminRepo
	authRepo      domain.IAuthRepo
}

// NewAdminService returns a new instance of the adminService struct.
func NewAdminService(adminRepo domain.IAdminRepo, authRepo domain.IAuthRepo) domain.IAdminService {
	return &adminService{
		requestHashes: make(map[string]bool),
		adminRepo:     adminRepo,
		authRepo:      authRepo,
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

// AddExecutiveCommitteeMember adds a new executive committee member.
func (adminService *adminService) AddExecutiveCommitteeMember(request *types.CreateCommitteeRequest) error {
	// Generate a hash of the request parameters
	doHash := sha256.New()
	doHash.Write([]byte(fmt.Sprintf("%s%s%s", request.Role, request.Name, request.Designation)))
	hash := hex.EncodeToString(doHash.Sum(nil))

	// Check if the hash exists in the map
	if _, exists := adminService.requestHashes[hash]; exists {
		// If it exists, return an error
		return fmt.Errorf("duplicate request")
	}
	// If it doesn't exist, add it to the map
	adminService.requestHashes[hash] = true

	// pass the request to the repository layer
	executiveCommitteeMember := &models.ExecutiveCommittee{}

	executiveCommitteeMember.Role = request.Role
	executiveCommitteeMember.Name = request.Name
	executiveCommitteeMember.Designation = request.Designation

	if err := adminService.adminRepo.AddExecutiveCommitteeMember(executiveCommitteeMember); err != nil {
		return err
	}
	return nil
}

func (adminService *adminService) DeleteExecutiveCommitteeMember(id string) error {
	// Get the executive committee member from the database
	execMember, err := adminService.adminRepo.FindExecutiveCommitteeMemberById(id)
	if err != nil {
		return err
	}

	// deleted member should be removed from the requestHashes map so that after deleting the member, the same member can be added again
	doHash := sha256.New()
	doHash.Write([]byte(fmt.Sprintf("%s%s%s", execMember.Role, execMember.Name, execMember.Designation)))
	hash := hex.EncodeToString(doHash.Sum(nil))

	// Remove the hash from the map
	delete(adminService.requestHashes, hash)

	// pass the request to the repository layer
	if err := adminService.adminRepo.DeleteExecutiveCommitteeMember(&execMember); err != nil {
		return err
	}
	return nil
}
