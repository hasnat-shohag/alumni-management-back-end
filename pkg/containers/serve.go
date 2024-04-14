package containers

import (
	"alumni-management-server/pkg/config"
	"alumni-management-server/pkg/connection"
	"alumni-management-server/pkg/controllers"
	"alumni-management-server/pkg/repositories"
	"alumni-management-server/pkg/routes"
	"alumni-management-server/pkg/services"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
)

func Serve(e *echo.Echo) {
	// config initialization
	config.SetConfig()

	// database connection
	connection.Connect()
	db := connection.GetDB()

	// repository initialization
	userRepository := repositories.UserDBInstance(db)
	adminRepository := repositories.AdminDBInstance(db)

	// service initialization
	userService := services.AuthServiceInstance(userRepository)
	adminService := services.NewAdminService(adminRepository)

	// controller initialization
	userController := controllers.NewAuthController(userService)
	adminController := controllers.NewAdminController(adminService)

	// route initialization
	userRoutes := routes.NewAuthRoutes(e, userController)
	adminRoutes := routes.NewAdminRoutes(e, adminController)
	userRoutes.InitAuthRoutes()
	adminRoutes.InitAdminRoutes()

	// starting server
	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))
}
