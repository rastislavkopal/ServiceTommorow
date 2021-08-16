package routes

import (
	"backend/api/controller"
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
	user := u.Handler.Gin.Group("/user")
	{
		user.GET("/", u.Controller.GetUsers)
		user.POST("/register", u.Controller.RegisterUser)
		user.POST("/login", u.Controller.LoginUser)
		// user.GET("/logout", u.Controller.LogoutUser)
		user.POST("/", u.Controller.AddUser)
		user.GET("/:id", u.Controller.GetUser)
		user.DELETE("/:id", u.Controller.DeleteUser)
		user.PUT("/:id", u.Controller.UpdateUser)
	}
}
