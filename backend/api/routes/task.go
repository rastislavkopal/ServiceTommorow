package routes

import (
	"backend/api/controller"
	"backend/common"
)

//TaskRoute -> Route for Task module
type TaskRoute struct {
	controller *controller.TaskController
	handler    *common.GinRouter
}

//NewTaskRoute -> initializes new tasks routes
func NewTaskRoute(c *controller.TaskController, h *common.GinRouter) TaskRoute {
	return TaskRoute{
		controller: c,
		handler:    h,
	}
}

//Setup -> setups new tasks Routes
func (t *TaskRoute) Setup() {

	task := t.handler.Gin.Group("/workspace/:ws_id/task")
	{
		task.GET("/", t.controller.GetTasks)
		task.POST("/", t.controller.AddTask)
		task.GET("/:id", t.controller.GetTask)
		task.DELETE("/:id", t.controller.DeleteTask)
		task.PUT("/:id", t.controller.UpdateTask)
	}

}
