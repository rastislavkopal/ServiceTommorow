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

	task := t.handler.Gin.Group("/workspace/:id/task")
	{
		task.GET("/", t.controller.GetTasks)
		task.POST("/", t.controller.AddTask)
		task.GET("/:task_id", t.controller.GetTask)
		task.DELETE("/:task_id", t.controller.DeleteTask)
		task.PUT("/:task_id", t.controller.UpdateTask)
	}

}
