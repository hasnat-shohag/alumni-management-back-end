package routes

import (
	"alumni-management-server/pkg/controllers"
	"alumni-management-server/pkg/middlewares"

	"github.com/labstack/echo/v4"
)

type ExecutiveCommitteeRoutes struct {
	echo                  *echo.Echo
	executiveCommitteeCtr controllers.ExecutiveCommitteeController
}

func NewExecutiveCommitteeRoutes(echo *echo.Echo, executiveCommitteeCtr *controllers.ExecutiveCommitteeController) *ExecutiveCommitteeRoutes {
	return &ExecutiveCommitteeRoutes{
		echo:                  echo,
		executiveCommitteeCtr: *executiveCommitteeCtr,
	}
}

func (routes *ExecutiveCommitteeRoutes) InitExecutiveCommitteeRoutes() {
	e := routes.echo
	v1 := e.Group("/v1")
	v1.GET("/executive-committee/list", routes.executiveCommitteeCtr.GetAllMember)

	v1.Use(middlewares.ValidateToken)

	v1.POST("/executive-committee/create", routes.executiveCommitteeCtr.Create)
	v1.PATCH("/executive-committee/update/:id", routes.executiveCommitteeCtr.Update)
	v1.DELETE("/executive-committee/delete/:id", routes.executiveCommitteeCtr.Delete)
}
