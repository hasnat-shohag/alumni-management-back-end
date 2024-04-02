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

	// service initialization
	userService := services.AuthServiceInstance(userRepository)

	// controller initialization
	userController := controllers.NewAuthController(userService)

	// route initialization
	userRoutes := routes.NewAuthRoutes(e, userController)
	userRoutes.InitAuthRoutes()

	// starting server
	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))
}
