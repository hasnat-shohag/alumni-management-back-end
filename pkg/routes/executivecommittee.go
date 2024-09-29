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

	executiveCommittee := v1.Group("/executive-committee")

	executiveCommittee.GET("/list", routes.executiveCommitteeCtr.GetAllMember)
	executiveCommittee.GET("/:id", routes.executiveCommitteeCtr.GetMemberById)

	executiveCommittee.Use(middlewares.ValidateToken)

	executiveCommittee.POST("/create", routes.executiveCommitteeCtr.Create)
	executiveCommittee.PATCH("/update/:id", routes.executiveCommitteeCtr.Update)
	executiveCommittee.DELETE("/delete/:id", routes.executiveCommitteeCtr.Delete)
}
