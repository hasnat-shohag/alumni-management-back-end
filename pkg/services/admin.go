package services

import (
	"alumni-management-server/pkg/common/response"
	"alumni-management-server/pkg/domain"
	"alumni-management-server/pkg/email"
	"alumni-management-server/pkg/models"
	"alumni-management-server/pkg/types"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
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

// AddExecutiveCommitteeMember adds a new executive committee member.
func (adminService *adminService) AddExecutiveCommitteeMember(request *types.CreateCommitteeRequest) error {
	// check if the same request is already added
	_, err := adminService.adminRepo.FindBy("email", request.Email)
	if err == nil {
		return fmt.Errorf("executive committee member already exists")
	}

	// Open the executive committee member image file
	file, err := request.Image.Open()
	if err != nil {
		return err
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	// Create a new file in the desired location
	dirPath := "./images/executive_committee_avatar"
	imagePath := filepath.Join(dirPath, request.Email+"_"+request.Image.Filename)

	// Create the directory if it doesn't exist
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			return err
		}
	}

	dst, err := os.Create(imagePath)
	if err != nil {
		return err
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			return
		}
	}(dst)

	// Copy the uploaded file to the new file
	if _, err := io.Copy(dst, file); err != nil {
		return err
	}

	// pass the request to the repository layer
	executiveCommitteeMember := &models.ExecutiveCommittee{}

	executiveCommitteeMember.Role = request.Role
	executiveCommitteeMember.Name = request.Name
	executiveCommitteeMember.Email = request.Email
	executiveCommitteeMember.Designation = request.Designation
	executiveCommitteeMember.ImagePath = imagePath

	if err := adminService.adminRepo.AddExecutiveCommitteeMember(executiveCommitteeMember); err != nil {
		return err
	}
	return nil
}

func (adminService *adminService) DeleteExecutiveCommitteeMember(id string) error {
	// Get the executive committee member from the database
	execMember, err := adminService.adminRepo.FindBy("id", id)
	if err != nil {
		return err
	}

	// pass the request to the repository layer
	if err := adminService.adminRepo.DeleteExecutiveCommitteeMember(&execMember); err != nil {
		return err
	}
	return nil
}

// UpdateExecutiveCommitteeMember updates an executive committee member.
func (adminService *adminService) UpdateExecutiveCommitteeMember(id string, request *types.UpdateCommitteeRequest) error {
	// Get the executive committee member from the database
	execMember, err := adminService.adminRepo.FindBy("id", id)
	if err != nil {
		return err
	}

	// Open the executive committee member image file
	file, err := request.Image.Open()
	if err != nil {
		return err
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	// Create a new file in the desired location
	dirPath := "./images/executive_committee_avatar"
	imagePath := filepath.Join(dirPath, request.Email+"_"+request.Image.Filename)

	// Create the directory if it doesn't exist
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			return err
		}
	}

	dst, err := os.Create(imagePath)
	if err != nil {
		return err
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			return
		}
	}(dst)

	// Copy the uploaded file to the new file
	if _, err := io.Copy(dst, file); err != nil {
		return err
	}

	// pass the request to the repository layer
	execMember.Role = request.Role
	execMember.Name = request.Name
	execMember.Email = request.Email
	execMember.Designation = request.Designation
	execMember.ImagePath = imagePath

	if err := adminService.adminRepo.UpdateExecutiveCommitteeMember(&execMember); err != nil {
		return err
	}
	return nil
}

// GetExecutiveCommitteeInfo returns the executive committee information.
func (adminService *adminService) GetExecutiveCommitteeInfo() ([]response.ExecutiveCommitteeResponse, error) {
	// pass the request to the repository layer
	executiveCommittee, err := adminService.adminRepo.GetExecutiveCommitteeInfo()
	if err != nil {
		return nil, err
	}

	var responses []response.ExecutiveCommitteeResponse
	for _, member := range executiveCommittee {
		responseMember := response.ExecutiveCommitteeResponse{
			ID:          strconv.Itoa(int(member.ID)),
			Role:        member.Role,
			Name:        member.Name,
			Designation: member.Designation,
		}
		responses = append(responses, responseMember)
	}

	return responses, nil

}
