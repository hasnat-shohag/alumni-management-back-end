package routes

import (
	"alumni-management-server/pkg/controllers"
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
}
