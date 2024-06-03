package routes

import (
	"alumni-management-server/pkg/controllers"
	"alumni-management-server/pkg/middlewares"
	"github.com/labstack/echo/v4"
)

type UserRoutes struct {
	echo    *echo.Echo
	UserCtr controllers.UserController
}

func NewUserRoutes(echo *echo.Echo, userCtr controllers.UserController) *UserRoutes {
	return &UserRoutes{
		echo:    echo,
		UserCtr: userCtr,
	}
}

func (routes *UserRoutes) InitUserRoutes() {
	e := routes.echo
	user := e.Group("/user")
	user.POST("/forget-password", routes.UserCtr.ForgetPassword)
	user.POST("/reset-password", routes.UserCtr.ResetPassword)
	user.GET("/ping", routes.UserCtr.Ping)

	user.Use(middlewares.ValidateToken)

	user.GET("/alumni-list", routes.UserCtr.GetAllAlumni)
	user.GET("/:id", routes.UserCtr.GetUser)
	user.DELETE("/delete-me/:id", routes.UserCtr.DeleteMe) // note [when user is deleted accesstoken still work!!
	user.PATCH("/update-me/:id", routes.UserCtr.UpdateMe)
}
