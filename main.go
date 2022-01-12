package main

import (
	"deployment/config"
	"deployment/delivery/controllers"
	"deployment/repository"
	"deployment/utils"
	"fmt"

	"deployment/delivery/routers"

	"github.com/labstack/echo/v4"
)

func main() {
	config := config.GetConfig()
	db := utils.InitDB(config)
	userRepo := repository.InstenceRepoUser(db)
	bookRepo := repository.NewRepositoryBook(db)
	controllerBook := controllers.NewControllersBook(bookRepo)
	controllerUser := controllers.InstenceControllerUser(userRepo)

	e := echo.New()
	routers.Users(e, controllerUser)
	routers.Books(e, controllerBook)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}
