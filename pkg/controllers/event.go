package controllers

import (
	"alumni-management-server/pkg/services"
	"github.com/labstack/echo/v4"
)

type EventControllerInterface interface {
	Create(context echo.Context) error
}

type EventController struct {
	eventService services.EventServiceInterface
}

func NewEventController(eventService services.EventServiceInterface) EventController {
	return EventController{eventService: eventService}
}

func (eventController *EventController) Create(context echo.Context) error {

	return nil
}
