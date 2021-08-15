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
	post := u.Handler.Gin.Group("/users")
	{
		post.GET("/", u.Controller.GetUsers)
		post.POST("/", u.Controller.AddUser)
		post.GET("/:id", u.Controller.GetUser)
		post.DELETE("/:id", u.Controller.DeleteUser)
		post.PUT("/:id", u.Controller.UpdateUser)
	}
}