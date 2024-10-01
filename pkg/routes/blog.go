package routes

import (
	"alumni-management-server/pkg/controllers"
	"alumni-management-server/pkg/middlewares"
	"github.com/labstack/echo/v4"
)

type BlogRoutes struct {
	echo    *echo.Echo
	blogCtr controllers.BlogController
}

func NewBlogRoutes(echo *echo.Echo, blogCtr *controllers.BlogController) *BlogRoutes {
	return &BlogRoutes{
		echo:    echo,
		blogCtr: *blogCtr,
	}
}

func (routes *BlogRoutes) InitBlogRoutes() {
	e := routes.echo
	v1 := e.Group("/v1")
	blog := v1.Group("/blog")

	blog.Use(middlewares.ValidateToken)

	blog.POST("/create", routes.blogCtr.Create)
	//blog.PATCH("/update/:id", routes.blogCtr.Update)
	//blog.DELETE("/delete/:id", routes.blogCtr.Delete)
	//blog.GET("/:id", routes.blogCtr.FindById)
	//blog.GET("/all", routes.blogCtr.FindAll)
}
