package main

import (
	"alumni-management-server/pkg/containers"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

// CORSConfig returns the CORS middleware configuration
func CORSConfig() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	})
}

func main() {
	e := echo.New()

	// Use CORS middleware
	e.Use(CORSConfig())

	// Other middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	containers.Serve(e)
}
