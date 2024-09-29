package routes

import (
	"alumni-management-server/pkg/controllers"
	"alumni-management-server/pkg/middlewares"
	"github.com/labstack/echo/v4"
)

type EventRoutes struct {
	echo     *echo.Echo
	eventCtr controllers.EventController
}

func NewEventRoutes(echo *echo.Echo, eventCtr *controllers.EventController) *EventRoutes {
	return &EventRoutes{
		echo:     echo,
		eventCtr: *eventCtr,
	}
}

func (routes *EventRoutes) InitEventRoutes() {
	e := routes.echo
	v1 := e.Group("/v1")

	event := v1.Group("/event")
	event.Use(middlewares.ValidateToken)

	event.POST("/create", routes.eventCtr.Create)
	event.PATCH("/update/:id", routes.eventCtr.Update)
	event.DELETE("/delete/:id", routes.eventCtr.Delete)
	event.GET("/:id", routes.eventCtr.FindById)
	event.GET("", routes.eventCtr.FindAll)
}
