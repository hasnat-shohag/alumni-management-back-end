package controllers

import (
	"alumni-management-server/pkg/common/logger"
	"alumni-management-server/pkg/common/response"
	"alumni-management-server/pkg/serializer"
	"alumni-management-server/pkg/services"
	"alumni-management-server/pkg/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type EventControllerInterface interface {
	Create(context echo.Context) error
	Update(context echo.Context) error
	Delete(context echo.Context) error
	FindById(context echo.Context) error
}

type EventController struct {
	eventService services.EventServiceInterface
}

func NewEventController(eventService services.EventServiceInterface) EventController {
	return EventController{eventService: eventService}
}

func (eventController *EventController) Create(context echo.Context) error {
	// Get the user role from the context
	role := context.Get("role").(string)
	if role != "admin" {
		return context.JSON(http.StatusForbidden, "only admins can add executive members")
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

	title := context.FormValue("title")
	eventDate := context.FormValue("event_date")
	startTime := context.FormValue("start_time")
	location := context.FormValue("location")
	description := context.FormValue("description")

	newEvent := serializer.CreateEventRequest{}

	newEvent.Image = fileHeader
	newEvent.Title = title
	newEvent.EventDate = eventDate
	newEvent.StartTime = startTime
	newEvent.Location = location
	newEvent.Description = description

	if err := context.Bind(&newEvent); err != nil {
		return context.JSON(response.GenerateErrorResponseBody(response.ErrParsingRequestBody))
	}

	//pass to the service layer
	event, err := eventController.eventService.Create(&newEvent)
	if err != nil {
		logger.Error(err)
		return context.JSON(response.GenerateErrorResponseBody(err))
	}

	return context.JSON(http.StatusCreated, response.GenerateSuccessResponse("created successfully", event.ID))
}

func (eventController *EventController) Update(context echo.Context) error {
	// Get the user role from the context
	role := context.Get("role").(string)
	if role != "admin" {
		return context.JSON(http.StatusForbidden, "only admins can add executive members")
	}

	// get value from the request body
	fileHeader, err := context.FormFile("image")
	if err != nil {
		return context.JSON(response.GenerateErrorResponseBody(response.ErrInvalidRequestParams))
	}

	// Check the file type
	if fileHeader.Header.Get("Content-Type") != "image/jpeg" && fileHeader.Header.Get("Content-Type") != "image/png" {
		return context.JSON(http.StatusBadRequest, "invalid file type: expected image")
	}

	eventID, err := utils.ParseParamAsInt(context, "id")
	title := context.FormValue("title")
	eventDate := context.FormValue("event_date")
	startTime := context.FormValue("start_time")
	location := context.FormValue("location")
	description := context.FormValue("description")

	newEvent := serializer.CreateEventRequest{}

	newEvent.Image = fileHeader
	newEvent.Title = title
	newEvent.EventDate = eventDate
	newEvent.StartTime = startTime
	newEvent.Location = location
	newEvent.Description = description

	if err := context.Bind(&newEvent); err != nil {
		return context.JSON(response.GenerateErrorResponseBody(response.ErrParsingRequestBody))
	}

	//pass to the service layer
	updatedEvent, err := eventController.eventService.Update(eventID, &newEvent)
	if err != nil {
		logger.Error(err)
		return context.JSON(response.GenerateErrorResponseBody(err))
	}

	return context.JSON(http.StatusOK, response.GenerateSuccessResponse("updated successfully", updatedEvent.ID))
}

func (eventController *EventController) Delete(context echo.Context) error {
	// Get the user role from the context
	role := context.Get("role").(string)
	if role != "admin" {
		return context.JSON(http.StatusForbidden, "only admins can add executive members")
	}

	eventID, err := utils.ParseParamAsInt(context, "id")
	if err != nil {
		return context.JSON(response.GenerateErrorResponseBody(response.ErrInvalidRequestParams))
	}

	ID, err := eventController.eventService.Delete(eventID)
	if err != nil {
		logger.Error(err)
		return context.JSON(response.GenerateErrorResponseBody(err))
	}

	return context.JSON(http.StatusOK, response.GenerateSuccessResponse("deleted successfully", ID))
}

func (eventController *EventController) FindById(context echo.Context) error {
	eventID, err := utils.ParseParamAsInt(context, "id")
	if err != nil {
		return context.JSON(response.GenerateErrorResponseBody(response.ErrInvalidRequestParams))
	}

	event, err := eventController.eventService.FindById(eventID)
	if err != nil {
		logger.Error(err)
		return context.JSON(response.GenerateErrorResponseBody(err))
	}

	return context.JSON(http.StatusOK, response.GenerateSuccessResponse("successful", event))
}
