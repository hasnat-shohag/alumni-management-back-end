package services

import (
	"alumni-management-server/pkg/common/response"
	"alumni-management-server/pkg/models"
	"alumni-management-server/pkg/repositories"
	"alumni-management-server/pkg/serializer"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
)

type ExecutiveCommitteeServiceInterface interface {
	Create(request *serializer.CreateCommitteeRequest) error
	Update(id string, request *serializer.UpdateCommitteeRequest) error
	Delete(id string) error
	GetAllMember() ([]response.ExecutiveCommitteeResponse, error)
}

type ExecutiveCommitteeService struct {
	executiveCommitteeRepo repositories.ExecutiveCommitteeRepoInterface
}

func NewExecutiveCommitteeService(executiveCommitteeRepo repositories.ExecutiveCommitteeRepoInterface) ExecutiveCommitteeService {
	return ExecutiveCommitteeService{executiveCommitteeRepo: executiveCommitteeRepo}
}

// Create a new executive committee member.
func (executiveCommitteeService *ExecutiveCommitteeService) Create(request *serializer.CreateCommitteeRequest) error {
	// check if the same request is already added
	_, err := executiveCommitteeService.executiveCommitteeRepo.FindBy("email", request.Email)
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

	if err := executiveCommitteeService.executiveCommitteeRepo.AddExecutiveCommitteeMember(executiveCommitteeMember); err != nil {
		return err
	}
	return nil
}

func (executiveCommitteeService *ExecutiveCommitteeService) Delete(id string) error {
	// Get the executive committee member from the database
	execMember, err := executiveCommitteeService.executiveCommitteeRepo.FindBy("id", id)
	if err != nil {
		return err
	}

	// pass the request to the repository layer
	if err := executiveCommitteeService.executiveCommitteeRepo.DeleteExecutiveCommitteeMember(&execMember); err != nil {
		return err
	}
	return nil
}

// UpdateExecutiveCommitteeMember updates an executive committee member.
func (executiveCommitteeService *ExecutiveCommitteeService) Update(id string, request *serializer.UpdateCommitteeRequest) error {
	// Get the executive committee member from the database
	execMember, err := executiveCommitteeService.executiveCommitteeRepo.FindBy("id", id)
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

	if err := executiveCommitteeService.executiveCommitteeRepo.UpdateExecutiveCommitteeMember(&execMember); err != nil {
		return err
	}
	return nil
}

// GetExecutiveCommitteeInfo returns the executive committee information.
func (executiveCommitteeService *ExecutiveCommitteeService) GetAllMember() ([]response.ExecutiveCommitteeResponse, error) {
	// pass the request to the repository layer
	executiveCommittee, err := executiveCommitteeService.executiveCommitteeRepo.GetExecutiveCommitteeInfo()
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
			Email:       member.Email,
			Image:       "http://localhost:9030/get-image/" + member.ImagePath,
		}
		responses = append(responses, responseMember)
	}

	return responses, nil

}
