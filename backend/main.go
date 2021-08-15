package main

import (
	"backend/api/controller"
	"backend/api/repository"
	"backend/api/routes"
	"backend/api/service"
	"backend/common"
	"backend/models"
)

func init() {
	// common.LoadEnv()
}

func main() {

	router := common.NewGinRouter()                             //router has been initialized and configured
	db := common.NewDatabase()                                  // databse has been initialized and configured
	userRepository := repository.NewUserRepository(db)          // repository are being setup
	userService := service.NewUserService(userRepository)       // service are being setup
	userController := controller.NewUserController(userService) // controller are being set up
	userRoute := routes.NewUserRoute(userController, router)    // user routes are initialized
	userRoute.Setup()                                           // user routes are being setup

	db.DB.AutoMigrate(&models.User{}) // migrating Post model to datbase table
	router.Gin.Run(":8000")           //server started on 8000 port
}
