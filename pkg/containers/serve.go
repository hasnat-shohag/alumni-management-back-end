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
	authRepository := repositories.AuthDBInstance(db)
	adminRepository := repositories.AdminDBInstance(db)
	userRepository := repositories.UserDBInstance(db)

	// service initialization
	authService := services.AuthServiceInstance(authRepository)
	adminService := services.NewAdminService(adminRepository)
	userService := services.UserServiceInstance(userRepository, authRepository)

	// controller initialization
	authController := controllers.NewAuthController(authService)
	adminController := controllers.NewAdminController(adminService)
	userController := controllers.NewUserController(userService)

	// route initialization
	authRoutes := routes.NewAuthRoutes(e, authController)
	adminRoutes := routes.NewAdminRoutes(e, adminController)
	userRoutes := routes.NewUserRoutes(e, userController)

	authRoutes.InitAuthRoutes()
	adminRoutes.InitAdminRoutes()
	userRoutes.InitUserRoutes()

	// starting server
	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))
}
