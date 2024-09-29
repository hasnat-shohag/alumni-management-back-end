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
	e.GET("/get-image/:image-path", routes.UserCtr.GetImage)

	v1 := e.Group("/v1")
	v1.GET("/ping", routes.UserCtr.Ping)

	user := v1.Group("/user")

	user.POST("/forget-password", routes.UserCtr.ForgetPassword)
	user.POST("/reset-password", routes.UserCtr.ResetPassword)

	user.Use(middlewares.ValidateToken)

	user.GET("/alumni-list", routes.UserCtr.GetAllAlumni)
	user.GET("/:id", routes.UserCtr.GetAlumni)
	user.DELETE("/delete-me/:id", routes.UserCtr.DeleteMe) // note [when user is deleted access token still work!!]
	user.PATCH("/complete-profile/:id", routes.UserCtr.UpdateMe)

	// Event
	//user.GET("/create-event", routes.UserCtr.CreateEvent)
}
