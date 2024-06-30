package controllers

import (
	"alumni-management-server/pkg/domain"
	"alumni-management-server/pkg/types"
	"github.com/labstack/echo/v4"
	"mime/multipart"
	"net/http"
)

// IAuthController is an interface that defines the methods implemented by the AuthController struct.
type IAuthController interface {
	Signup(e echo.Context) error
	Login(e echo.Context) error
}

// AuthController defines the methods of the IAuthController interface.
type AuthController struct {
	authSvc domain.IAuthService
}

// NewAuthController is a function that returns a new instance of the AuthController struct.
func NewAuthController(authSvc domain.IAuthService) AuthController {
	return AuthController{
		authSvc: authSvc,
	}
}

func (authController *AuthController) Signup(context echo.Context) error {

	// Get the image file from the form data
	fileHeader, err := context.FormFile("certificate_or_id_card")
	if err != nil {
		return context.JSON(http.StatusBadRequest, "invalid image file")
	}

	// Check the file type
	if fileHeader.Header.Get("Content-Type") != "image/jpeg" && fileHeader.Header.Get("Content-Type") != "image/png" {
		return context.JSON(http.StatusBadRequest, "invalid file type: expected image")
	}

	name := context.FormValue("name")
	studentId := context.FormValue("student_id")
	email := context.FormValue("email")
	graduationYear := context.FormValue("graduation_year")
	role := context.FormValue("role")
	password := context.FormValue("password")
	confirmPassword := context.FormValue("confirm_password")

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

	// bind the request body to the SignupRequest struct
	registerRequest := &types.SignupRequest{
		Name:                name,
		StudentId:           studentId,
		Email:               email,
		GraduationYear:      graduationYear,
		Role:                role,
		CertificateOrIdCard: fileHeader,
		Password:            password,
		ConfirmPassword:     confirmPassword,
	}

	if err := context.Bind(registerRequest); err != nil {
		return context.JSON(http.StatusBadRequest, "invalid request body")
	}

	// validate the request body
	if err := registerRequest.Validate(); err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	// pass the request to the service layer
	if err := authController.authSvc.SignupUser(registerRequest); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusCreated, "user was created successfully")
}

func (authController *AuthController) Login(e echo.Context) error {
	loginRequest := &types.LoginRequest{}
	if err := e.Bind(loginRequest); err != nil {
		return e.JSON(http.StatusBadRequest, "invalid request body")
	}

	if err := loginRequest.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	loginResponse, err := authController.authSvc.Login(loginRequest)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, loginResponse)
}
