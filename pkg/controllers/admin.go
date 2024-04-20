package controllers

import (
	"alumni-management-server/pkg/domain"
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

	// pass the request to the service layer
	if err := adminController.AdminSvc.VerifyUser(studentId, isValid); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "User verified successfully")
}