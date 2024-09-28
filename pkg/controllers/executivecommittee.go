package controllers

import (
	"alumni-management-server/pkg/common/response"
	"alumni-management-server/pkg/serializer"
	"alumni-management-server/pkg/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ExecutiveCommitteeControllerInterface interface {
	Create(context echo.Context) error
	Update(context echo.Context) error
	Delete(context echo.Context) error
	GetAllMember(context echo.Context) ([]response.ExecutiveCommitteeResponse, error)
	GetMemberById(context echo.Context) (response.ExecutiveCommitteeResponse, error)
}

type ExecutiveCommitteeController struct {
	executiveCommitteeService services.ExecutiveCommitteeServiceInterface
}

func NewExecutiveCommitteeController(executiveCommitteeService services.ExecutiveCommitteeServiceInterface) ExecutiveCommitteeController {
	return ExecutiveCommitteeController{executiveCommitteeService: executiveCommitteeService}
}

// create a new executive committee member.
func (executiveCommitteeController *ExecutiveCommitteeController) Create(context echo.Context) error {
	// Get the user role from the context
	role := context.Get("role").(string)
	if role != "admin" {
		return context.JSON(http.StatusForbidden, "only admins can add executive members")
	}

	// Get the image from the request body
	fileHeader, err := context.FormFile("image")
	if err != nil {
		return context.JSON(http.StatusBadRequest, "invalid image file")
	}

	// Check the file type
	if fileHeader.Header.Get("Content-Type") != "image/jpeg" && fileHeader.Header.Get("Content-Type") != "image/png" {
		return context.JSON(http.StatusBadRequest, "invalid file type: expected png/jpg/jpg image")
	}

	role = context.FormValue("role")
	name := context.FormValue("name")
	email := context.FormValue("email")
	designation := context.FormValue("designation")

	// bind the request body to the CreateCommitteeRequest struct
	createCommitteeRequest := &serializer.CreateCommitteeRequest{
		Role:        role,
		Name:        name,
		Email:       email,
		Designation: designation,
		Image:       fileHeader,
	}

	if err := context.Bind(createCommitteeRequest); err != nil {
		return context.JSON(http.StatusBadRequest, "invalid request body")
	}

	// validate the request body
	if err := createCommitteeRequest.Validate(); err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	// pass the request to the service layer
	if err := executiveCommitteeController.executiveCommitteeService.Create(createCommitteeRequest); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusCreated, "executive committee member added successfully")
}

// delete an executive committee member
func (executiveCommitteeController *ExecutiveCommitteeController) Delete(context echo.Context) error {
	id := context.Param("id")

	// Get the user role from the context
	role := context.Get("role").(string)
	if role != "admin" {
		return context.JSON(http.StatusForbidden, "only admins can delete executive members")
	}

	// pass the request to the service layer
	if err := executiveCommitteeController.executiveCommitteeService.Delete(id); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, "executive committee member deleted successfully")
}

// UpdateExecutiveCommitteeMember updates an executive committee member
func (executiveCommitteeController *ExecutiveCommitteeController) Update(context echo.Context) error {
	id := context.Param("id")

	// Get the user role from the context
	roleFromToken := context.Get("role").(string)
	if roleFromToken != "admin" {
		return context.JSON(http.StatusForbidden, "only admins can update executive members")
	}

	// get the form values
	role := context.FormValue("role")
	name := context.FormValue("name")
	email := context.FormValue("email")
	designation := context.FormValue("designation")
	fileHeader, err := context.FormFile("image")
	if err != nil {
		return context.JSON(http.StatusBadRequest, "invalid image file")
	}

	// Check the file type
	if fileHeader.Header.Get("Content-Type") != "image/jpeg" && fileHeader.Header.Get("Content-Type") != "image/png" {
		return context.JSON(http.StatusBadRequest, "invalid file type: expected png/jpg/jpg image")
	}

	updateCommitteeRequest := &serializer.UpdateCommitteeRequest{}
	updateCommitteeRequest.Role = role
	updateCommitteeRequest.Name = name
	updateCommitteeRequest.Email = email
	updateCommitteeRequest.Designation = designation
	updateCommitteeRequest.Image = fileHeader

	if err := context.Bind(updateCommitteeRequest); err != nil {
		return context.JSON(http.StatusBadRequest, "invalid request body")
	}

	// validate the request body
	if err := updateCommitteeRequest.Validate(); err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	// pass the request to the service layer
	if err := executiveCommitteeController.executiveCommitteeService.Update(id, updateCommitteeRequest); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, "executive committee member updated successfully")
}

// returns the executive all the committee member information.
func (executiveCommitteeController *ExecutiveCommitteeController) GetAllMember(context echo.Context) error {
	// pass request to the service layer
	committeeInfo, err := executiveCommitteeController.executiveCommitteeService.GetAllMember()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, committeeInfo)
}

// returns the executive committee member information by id.
func (executiveCommitteeController *ExecutiveCommitteeController) GetMemberById(context echo.Context) error {
	id := context.Param("id")

	// pass request to the service layer
	committeeInfo, err := executiveCommitteeController.executiveCommitteeService.GetMemberById(id)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, committeeInfo)
}
