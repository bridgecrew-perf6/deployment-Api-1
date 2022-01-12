package routers

import (
	"deployment/config"
	"deployment/delivery/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Users(e *echo.Echo, uc *controllers.StructControllerUser) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, path=${path} status=${status}\n",
	}))

	aut := e.Group("/auth")
	aut.Use(middleware.JWT([]byte(config.JWT_SECRET)))
	//================
	//ROUTER FOR USERS
	//================
	aut.GET("/users", uc.GetUsersController)
	aut.GET("/users/:id", uc.GetUserController)
	e.POST("/users", uc.CreateUserController)
	aut.DELETE("/users/:id", uc.DeleteUserController)
	aut.PUT("/users/:id", uc.UpdateUserController)
	e.POST("/login", uc.Login)
}

func Books(e *echo.Echo, uc *controllers.Controllers) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, path=${path} status=${status}\n",
	}))

	aut := e.Group("/auth")
	aut.Use(middleware.JWT([]byte(config.JWT_SECRET)))
	//=================
	// ROUTER FOR BOOKS
	//=================
	e.GET("/books", uc.GetBooks)
	e.GET("/books/:id", uc.GetBookById)
	aut.POST("/books", uc.CreateBook)
	aut.DELETE("/books/:id", uc.DeleteBook)
	aut.PUT("/books/:id", uc.UpdateBook)

}
