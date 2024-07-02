package controllers

import (
	"alumni-management-server/pkg/common/response"
	"alumni-management-server/pkg/domain"
	"alumni-management-server/pkg/types"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// IAdminController is an interface that defines the methods implemented by the AdminController struct.
type IAdminController interface {
	VerifyUser(e echo.Context) error
}

// AdminController defines the methods of the IAdminController interface.
type AdminController struct {
	AdminSvc domain.IAdminService
}

func NewAdminController(adminSvc domain.IAdminService) AdminController {
	return AdminController{
		AdminSvc: adminSvc,
	}
}

func (adminController *AdminController) VerifyUser(c echo.Context) error {
	studentId := c.QueryParam("student_id")
	isValidSting := c.QueryParam("is_valid")
	isValid, err := strconv.ParseBool(isValidSting)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid value for is_valid")
	}

	// Get the user role from the context
	role := c.Get("role").(string)
	if role != "admin" {
		return c.JSON(http.StatusForbidden, "Only admins can verify users")
	}

	// pass the request to the service layer
	if err := adminController.AdminSvc.VerifyUser(studentId, isValid); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "User verified successfully")
}

func (adminController *AdminController) DeleteUser(c echo.Context) error {
	studentId := c.Param("id")

	// Get the user role from the context
	role := c.Get("role").(string)
	if role != "admin" {
		return c.JSON(http.StatusForbidden, "Only admins can delete users")
	}

	// pass the request to the service layer
	if err := adminController.AdminSvc.DeleteUser(studentId); err != nil {
		return c.JSON(response.GenerateErrorResponseBody(err))
	}
	return c.JSON(http.StatusOK, "User deleted successfully")
}

// AddExecutiveCommittee adds a new executive committee member.
func (adminController *AdminController) AddExecutiveCommittee(context echo.Context) error {
	createCommitteeRequest := &types.CreateCommitteeRequest{}
	if err := context.Bind(createCommitteeRequest); err != nil {
		return context.JSON(http.StatusBadRequest, "invalid request body")
	}

	// validate the request body
	if err := createCommitteeRequest.Validate(); err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	// Get the user role from the context
	role := context.Get("role").(string)
	if role != "admin" {
		return context.JSON(http.StatusForbidden, "only admins can add executive members")
	}

	// pass the request to the service layer
	if err := adminController.AdminSvc.AddExecutiveCommitteeMember(createCommitteeRequest); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusCreated, "executive committee member added successfully")
}

// DeleteExecutiveCommitteeMember delete an executive committee member
func (adminController *AdminController) DeleteExecutiveCommitteeMember(context echo.Context) error {
	id := context.Param("id")

	// Get the user role from the context
	role := context.Get("role").(string)
	if role != "admin" {
		return context.JSON(http.StatusForbidden, "only admins can delete executive members")
	}

	// pass the request to the service layer
	if err := adminController.AdminSvc.DeleteExecutiveCommitteeMember(id); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, "executive committee member deleted successfully")
}

// UpdateExecutiveCommitteeMember updates an executive committee member
func (adminController *AdminController) UpdateExecutiveCommitteeMember(context echo.Context) error {
	id := context.Param("id")

	// Get the user role from the context
	roleFromToken := context.Get("role").(string)
	if roleFromToken != "admin" {
		return context.JSON(http.StatusForbidden, "only admins can update executive members")
	}

	updateCommitteeRequest := &types.UpdateCommitteeRequest{}

	if err := context.Bind(updateCommitteeRequest); err != nil {
		return context.JSON(http.StatusBadRequest, "invalid request body")
	}

	// validate the request body
	if err := updateCommitteeRequest.Validate(); err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	// pass the request to the service layer
	if err := adminController.AdminSvc.UpdateExecutiveCommitteeMember(id, updateCommitteeRequest); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, "executive committee member updated successfully")
}
