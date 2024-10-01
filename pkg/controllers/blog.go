package controllers

import (
	"alumni-management-server/pkg/common/logger"
	"alumni-management-server/pkg/common/response"
	"alumni-management-server/pkg/serializer"
	"alumni-management-server/pkg/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type BlogControllerInterface interface {
	Create(context echo.Context) error
	Update(context echo.Context) error
	Delete(context echo.Context) error
	FindById(context echo.Context) error
	FindAll(context echo.Context) error
}

type BlogController struct {
	blogService services.BlogServiceInterface
}

func NewBlogController(blogService services.BlogServiceInterface) BlogController {
	return BlogController{blogService: blogService}
}

func (blogController *BlogController) Create(context echo.Context) error {
	// Get the user role from the context
	role := context.Get("role").(string)
	if role != "alumni" && role != "admin" {
		return context.JSON(http.StatusForbidden, "only alumni and admin can create blog posts")
	}

	// get value from the request body
	fileHeader, err := context.FormFile("image")
	if err != nil {
		return context.JSON(response.GenerateErrorResponseBody(response.ErrParsingRequestBody))
	}

	// Check the file type
	if fileHeader.Header.Get("Content-Type") != "image/jpeg" && fileHeader.Header.Get("Content-Type") != "image/png" {
		return context.JSON(http.StatusBadRequest, "invalid file type: expected image")
	}

	// get the user id from the context
	userId_, ok := context.Get("student_id").(string)
	if !ok {
		logger.Error("user_id not found in context or is of incorrect type ", userId_)
		return context.JSON(http.StatusInternalServerError, "internal server error")
	}

	// Convert userId from string to uint
	userId, err := strconv.ParseUint(userId_, 10, 32)
	if err != nil {
		logger.Error("failed to convert user_id to uint: ", err)
		return context.JSON(http.StatusInternalServerError, "internal server error")
	}
	// create a new post
	newPost := serializer.CreateBlogRequest{
		Image:    fileHeader,
		UserId:   uint(userId),
		Title:    context.FormValue("title"),
		Content:  context.FormValue("content"),
		Category: context.FormValue("category"),
		Tags:     context.FormValue("tags"),
		Status:   context.FormValue("status"),
	}

	if err := context.Bind(&newPost); err != nil {
		return context.JSON(response.GenerateErrorResponseBody(response.ErrParsingRequestBody))
	}

	// pass to the service layer
	blog, err := blogController.blogService.Create(&newPost)
	if err != nil {
		logger.Error(err)
		return context.JSON(response.GenerateErrorResponseBody(err))
	}

	return context.JSON(http.StatusCreated, response.GenerateSuccessResponse("created successfully", blog.ID))
}
