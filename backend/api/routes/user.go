package routes

import (
	"backend/api/controller"
	"backend/api/middleware"
	"backend/common"
)

//UserRoute -> Route for question module
type UserRoute struct {
	controller *controller.UserController
	handler    *common.GinRouter
}

//NewUserRoute -> initializes new choice routes
func NewUserRoute(c *controller.UserController, h *common.GinRouter) UserRoute {
	return UserRoute{
		controller: c,
		handler:    h,
	}
}

//Setup -> setups new choice Routes
func (u *UserRoute) Setup() {

	auth := u.handler.Gin.Group("/auth")
	{
		auth.POST("/register", u.controller.RegisterUser)
		auth.POST("/login", u.controller.LoginUser)
		// user.GET("/logout", u.Controller.LogoutUser)
	}

	user := u.handler.Gin.Group("/user")
	user.Use(middleware.AuthRequired)
	{
		user.GET("/", u.controller.GetUsers)
		user.POST("/", u.controller.AddUser)
		user.GET("/:id", u.controller.GetUser)
		user.DELETE("/:id", u.controller.DeleteUser)
		user.PUT("/:id", u.controller.UpdateUser)
	}

}
