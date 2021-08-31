package controller

import (
	"backend/api/service"
	"backend/models"
	"backend/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//TaskController -> TaskController
type TaskController struct {
	service *service.TaskService
}

//NewTaskController : NewTaskController
func NewTaskController(s *service.TaskService) TaskController {
	return TaskController{
		service: s,
	}
}

// GetTasks : GetTasks controller
// @Summary Get list of tasks
// @Description get all tasks
// @Tags Tasks
// @Success 200 {array} models.Task
// @Failure 404 {object} object
// @Router / [get]
func (t *TaskController) GetTasks(ctx *gin.Context) {
	var tasks models.Task

	keyword := ctx.Query("title")

	data, total, err := t.service.FindAll(tasks, keyword)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find tasks")
		return
	}
	respArr := make([]map[string]interface{}, 0, 0)

	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Task result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

// AddTask : AddTask controller
func (t *TaskController) AddTask(ctx *gin.Context) {
	var task models.Task
	ctx.ShouldBindJSON(&task)

	if task.Title == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "FirstName is required")
		return
	}

	err := t.service.Save(task)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create task")
		return
	}
	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created User")
}

// GetUGetTaskser -> get task by id
// @Summary Get one task
// @Description get task by ID
// @Tags Tasks
// @Param id path string true "Task ID"
// @Success 200 {object} models.Task
// @Failure 400,404 {object} object
// @Router /{id} [get]
func (t *TaskController) GetTask(c *gin.Context) {
	idParam := c.Param("task_id")
	id, err := strconv.ParseUint(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	var task models.Task
	task.ID = id
	foundTask, err := t.service.Find(task)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Error Finding User")
		return
	}
	response := foundTask.ResponseMap()

	c.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of Task",
		Data:    &response})

}

//DeleteTask : Deletes task
func (t *TaskController) DeleteTask(c *gin.Context) {
	idParam := c.Param("task_id")
	id, err := strconv.ParseUint(idParam, 10, 64) //type conversion string to uint64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	err = t.service.Delete(id)

	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Failed to delete Task")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Deleted Sucessfully"}
	c.JSON(http.StatusOK, response)
}

//UpdateTask : update by id
func (t *TaskController) UpdateTask(ctx *gin.Context) {
	idParam := ctx.Param("task_id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var task models.Task
	task.ID = id

	taskRecord, err := t.service.Find(task)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Task with given id not found")
		return
	}
	ctx.ShouldBindJSON(&taskRecord)

	if err := t.service.Update(taskRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store Task")
		return
	}
	response := taskRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated Task",
		Data:    response,
	})
}
