package containers

import (
	"alumni-management-server/pkg/common/logger"
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
	// init logger
	logger.NewLogger()

	// config initialization
	config.SetConfig()

	// database connection
	connection.Connect()
	db := connection.GetDB()

	// event module initialization
	eventRepository := repositories.NewEventRepo(db)
	eventService := services.NewEventService(&eventRepository)
	eventController := controllers.NewEventController(&eventService)

	// Executive Committee module initialization
	executiveCommitteeRepository := repositories.NewExecutiveCommitteeRepo(db)
	executiveCommitteeService := services.NewExecutiveCommitteeService(&executiveCommitteeRepository)
	executiveCommitteeController := controllers.NewExecutiveCommitteeController(&executiveCommitteeService)

	// Blog module initialization
	blogRepository := repositories.NewBlogRepo(db)
	blogService := services.NewBlogService(&blogRepository)
	blogController := controllers.NewBlogController(&blogService)

	// repository initialization
	authRepository := repositories.AuthDBInstance(db)
	adminRepository := repositories.AdminDBInstance(db)
	userRepository := repositories.UserDBInstance(db)

	// service initialization
	authService := services.AuthServiceInstance(authRepository)
	adminService := services.NewAdminService(adminRepository, authRepository)
	userService := services.UserServiceInstance(userRepository, authRepository, adminRepository)

	// controller initialization
	authController := controllers.NewAuthController(authService)
	adminController := controllers.NewAdminController(adminService)
	userController := controllers.NewUserController(userService)

	// route initialization
	authRoutes := routes.NewAuthRoutes(e, authController)
	adminRoutes := routes.NewAdminRoutes(e, adminController)
	userRoutes := routes.NewUserRoutes(e, userController)
	eventRoutes := routes.NewEventRoutes(e, &eventController)
	executiveCommitteeRoutes := routes.NewExecutiveCommitteeRoutes(e, &executiveCommitteeController)
	blogRoutes := routes.NewBlogRoutes(e, &blogController)

	authRoutes.InitAuthRoutes()
	adminRoutes.InitAdminRoutes()
	userRoutes.InitUserRoutes()
	eventRoutes.InitEventRoutes()
	executiveCommitteeRoutes.InitExecutiveCommitteeRoutes()
	blogRoutes.InitBlogRoutes()

	// starting server
	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))
}
