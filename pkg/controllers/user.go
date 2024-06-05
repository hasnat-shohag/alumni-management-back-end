package controllers

import (
	"alumni-management-server/pkg/common/response"
	"alumni-management-server/pkg/domain"
	"alumni-management-server/pkg/types"
	"github.com/labstack/echo/v4"
	"mime/multipart"
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

	res := map[string]interface{}{
		"data": alumni,
		"meta": map[string]interface{}{
			"page":         page,
			"limit":        limit,
			"totalRecords": totalRecords,
			"totalPages":   totalPages,
		},
	}

	return context.JSON(http.StatusOK, res)
}

func (userController *UserController) GetUser(context echo.Context) error {
	studentId := context.Param("id")

	user, err := userController.userSvc.GetUser(studentId)
	if err != nil {
		return context.JSON(response.GenerateErrorResponseBody(err))
	}

	return context.JSON(http.StatusOK, user)
}

func (userController *UserController) DeleteMe(context echo.Context) error {

	studentId := context.Param("id")
	studentIdFromToken := context.Get("student_id").(string)

	err := userController.userSvc.DeleteMe(studentId, studentIdFromToken)
	if err != nil {
		return context.JSON(response.GenerateErrorResponseBody(err))
	}

	return context.JSON(http.StatusOK, "your account deleted successfully.")
}

func (userController *UserController) UpdateMe(context echo.Context) error {
	studentId := context.Param("id")
	studentIdFromToken := context.Get("student_id").(string)

	if studentId != studentIdFromToken {
		return context.JSON(http.StatusUnauthorized, "you have no access to update others account")
	}

	// Get the image file from the form data
	fileHeader, err := context.FormFile("image")
	if err != nil {
		return context.JSON(http.StatusBadRequest, "invalid image file")
	}

	// Check the file type
	if fileHeader.Header.Get("Content-Type") != "application/image" {
		return context.JSON(http.StatusBadRequest, "invalid file type: expected image")
	}

	jobType := context.FormValue("job_type")
	instituteName := context.FormValue("institute_name")
	jobTitle := context.FormValue("job_title")
	phoneNumber := context.FormValue("phone_number")
	linkedIn := context.FormValue("linked_in")
	facebook := context.FormValue("facebook")

	// Open the image file
	file, err := fileHeader.Open()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "unable to open image file")
	}

	// Close the image file after the function returns
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			context.Logger().Error(err)
		}
	}(file)

	completeProfileRequest := types.CompleteProfileRequest{
		Image:         fileHeader,
		JobType:       jobType,
		InstituteName: instituteName,
		JobTitle:      jobTitle,
		PhoneNumber:   phoneNumber,
		LinkedIn:      linkedIn,
		Facebook:      facebook,
	}

	if err := completeProfileRequest.Validate(); err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	err = userController.userSvc.UpdateMe(studentId, completeProfileRequest)
	if err != nil {
		return context.JSON(response.GenerateErrorResponseBody(err))
	}

	return context.JSON(http.StatusOK, "your account updated successfully.")
}

func (userController *UserController) GetImage(context echo.Context) error {
	imagePath := context.Param("image-path")

	return context.File("./" + imagePath)
}
