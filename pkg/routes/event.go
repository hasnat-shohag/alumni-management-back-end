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

	v1.Use(middlewares.ValidateToken)
	v1.POST("/event/create", routes.eventCtr.Create)
	v1.PATCH("/event/update/:id", routes.eventCtr.Update)
	v1.DELETE("/event/delete/:id", routes.eventCtr.Delete)
	v1.GET("/event/:id", routes.eventCtr.FindById)
}
