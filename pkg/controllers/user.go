package controllers

import (
	"alumni-management-server/pkg/common/response"
	"alumni-management-server/pkg/domain"
	"alumni-management-server/pkg/types"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserControllerInterface interface {
	ForgetPassword(context echo.Context) error
	ResetPasswordRequest(context echo.Context) error
	GetAllAlumni(context echo.Context) error
}

type UserController struct {
	userSvc domain.IUserService
}

func NewUserController(userSvc domain.IUserService) UserController {
	return UserController{
		userSvc: userSvc,
	}
}

func (userController *UserController) ForgetPassword(e echo.Context) error {
	forgotPasswordReq := types.ForgotPasswordRequest{}
	// bind the request body to the ForgotPasswordRequest struct
	if err := e.Bind(&forgotPasswordReq); err != nil {
		return e.JSON(http.StatusBadRequest, "invalid request body")
	}
	// validate the request body
	if err := forgotPasswordReq.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	// pass the request to the service layer
	if err := userController.userSvc.ForgetPassword(forgotPasswordReq.Email); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "Password reset instructions sent to your email. Check your inbox.")
}

func (userController *UserController) ResetPassword(e echo.Context) error {
	resetPasswordReq := types.ResetPasswordRequest{}

	// bind the request body to the ResetPasswordRequest struct
	if err := e.Bind(&resetPasswordReq); err != nil {
		return e.JSON(http.StatusBadRequest, "invalid request body")
	}
	// validate the request body
	if err := resetPasswordReq.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	otp := e.QueryParam("otp")
	email := e.QueryParam("email")

	// pass the request to the service layer
	user, err := userController.userSvc.GetUserFromEmailWithValidOtp(email, otp)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	// pass the request to the service layer
	if err := userController.userSvc.ResetPassword(user, resetPasswordReq.NewPassword); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "Password reset successfully.")
}

// Ping is a simple function that returns a message to the client.
func (userController *UserController) Ping(e echo.Context) error {
	return e.JSON(http.StatusOK, "pong")
}

func (userController *UserController) GetAllAlumni(context echo.Context) error {
	// get the query params
	page := context.QueryParam("page")
	limit := context.QueryParam("limit")
	// convert the query params to integers
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return context.JSON(response.GenerateErrorResponseBody(err))
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return context.JSON(response.GenerateErrorResponseBody(err))
	}

	if pageInt < 1 {
		pageInt = 1
	}
	if limitInt < 1 {
		limitInt = 10
	}

	// pass the request to the service layer
	alumni, totalRecords, err := userController.userSvc.GetAllAlumni(pageInt, limitInt)
	if err != nil {
		return context.JSON(response.GenerateErrorResponseBody(err))
	}

	totalPages := (totalRecords + limitInt - 1) / limitInt

	response := map[string]interface{}{
		"data": alumni,
		"meta": map[string]interface{}{
			"page":         page,
			"limit":        limit,
			"totalRecords": totalRecords,
			"totalPages":   totalPages,
		},
	}

	return context.JSON(http.StatusOK, response)
}
