package main

import (
	"backend/api/controller"
	"backend/api/repository"
	"backend/api/routes"
	"backend/api/service"
	"backend/common"
	"backend/models"
	"os"
)

func init() {
	// not needed when loaded env vars from docker-compose
	// common.LoadEnv()
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}
func main() {

	router := common.NewGinRouter()                             //router has been initialized and configured
	db := common.NewDatabase()                                  // databse has been initialized and configured
	userRepository := repository.NewUserRepository(db)          // repository are being setup
	userService := service.NewUserService(userRepository)       // service are being setup
	userController := controller.NewUserController(userService) // controller are being set up
	userRoute := routes.NewUserRoute(userController, router)    // user routes are initialized
	userRoute.Setup()                                           // user routes are being setup

	// migrating models to datbase table
	db.DB.AutoMigrate(
		&models.User{},
		&models.TokenDetails{},
	)

	router.Gin.Run(":" + os.Getenv("SERVER_PORT")) //server started on 8000 port
}
