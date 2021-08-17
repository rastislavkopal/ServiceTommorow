package routes

import (
	"backend/api/controller"
	"backend/api/middleware"
	"backend/common"
)

//UserRoute -> Route for question module
type UserRoute struct {
	Controller controller.UserController
	Handler    common.GinRouter
}

//NewUserRoute -> initializes new choice routes
func NewUserRoute(
	controller controller.UserController,
	handler common.GinRouter,
) UserRoute {
	return UserRoute{
		Controller: controller,
		Handler:    handler,
	}
}

//Setup -> setups new choice Routes
func (u UserRoute) Setup() {

	auth := u.Handler.Gin.Group("/auth")
	{
		auth.POST("/register", u.Controller.RegisterUser)
		auth.POST("/login", u.Controller.LoginUser)
		// user.GET("/logout", u.Controller.LogoutUser)
	}

	user := u.Handler.Gin.Group("/user")
	user.Use(middleware.AuthRequired)
	{
		user.GET("/", u.Controller.GetUsers)
		user.POST("/", u.Controller.AddUser)
		user.GET("/:id", u.Controller.GetUser)
		user.DELETE("/:id", u.Controller.DeleteUser)
		user.PUT("/:id", u.Controller.UpdateUser)
	}

}
