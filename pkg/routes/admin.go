package routes

import (
	"alumni-management-server/pkg/controllers"
	"github.com/labstack/echo/v4"
)

// AdminRoutes stores controller and echo instance for admin.
type AdminRoutes struct {
	echo     *echo.Echo
	adminCtr controllers.AdminController
}

// NewAdminRoutes returns a new instance of the AdminRoutes struct.
func NewAdminRoutes(echo *echo.Echo, adminCtr controllers.AdminController) *AdminRoutes {
	return &AdminRoutes{
		echo:     echo,
		adminCtr: adminCtr,
	}
}

// InitAdminRoutes initializes the admin routes.
func (routes *AdminRoutes) InitAdminRoutes() {
	e := routes.echo
	admin := e.Group("/admin")
	admin.POST("/verify-user/", routes.adminCtr.VerifyUser)
}
