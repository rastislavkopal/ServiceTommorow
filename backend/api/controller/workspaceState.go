package controller

import (
	"backend/api/service"
	"backend/models"
	"backend/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WorkspaceStateController struct {
	service *service.WorkspaceService
}

func NewWorkspaceStateController(s *service.WorkspaceService) WorkspaceStateController {
	return WorkspaceStateController{
		service: s,
	}
}

// @Summary Get list of workspaceStates
// @Description get all WorkspacesStates
// @Tags WorkspaceState
// @Success 200 {array} models.WorkspaceState
// @Failure 404 {object} object
// @Router / [get]
func (w *WorkspaceController) GetWorkspaceStates(ctx *gin.Context) {
	var wss models.WorkspaceState

	data, total, err := w.service.FindAllWorkspaceState(wss)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find WorkspaceState")
		return
	}
	respArr := make([]map[string]interface{}, 0, 0)

	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Workspace result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

// @Summary Create new WorkspaceState
// @Description Create new WorkspaceState
// @Tags WorkspaceState
// @Success 201 {array} models.WorkspaceState
// @Failure 400 {object} object
// @Router / [post]
func (w *WorkspaceController) CreateWorkspaceState(ctx *gin.Context) {
	var ws models.Workspace
	ctx.ShouldBindJSON(&ws)

	if ws.Title == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
		return
	}

	err := w.service.Save(&ws, ws.AuthorID)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create workspace")
		return
	}

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully created new workspace",
		Data: map[string]interface{}{
			"workspace_id": ws.ID,
		}})
}

// @Summary Get one WorkspaceState
// @Description get WorkspaceState by ID
// @Tags WorkspaceState
// @Param id path string true "WorkspaceState ID"
// @Success 200 {object} models.WorkspaceState
// @Failure 400,404 {object} object
// @Router /{id} [get]
func (w *WorkspaceController) GetWorkspaceState(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	workspace := models.Workspace{ID: id}
	foundWorkspace, err := w.service.Find(&workspace)

	if err != nil {
		util.ErrorJSON(c, http.StatusNotFound, "Workspace not found")
		return
	}
	response := foundWorkspace.ResponseMap()

	c.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result workspace",
		Data:    &response})
}

// @Summary Delete workspace by it's ID
// @Description Find WorkspaceState by queryID and delete it
// @Tags WorkspaceState
// @Param id path string true "WorkspaceState ID"
// @Success 200 {object} models.WorkspaceState
// @Failure 400,404 {object} object
// @Router /{id} [delete]
func (w *WorkspaceController) DeleteWorkspaceState(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64) //type conversion string to uint64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}

	err = w.service.Delete(id)

	if err != nil {
		util.ErrorJSON(c, http.StatusNotFound, "Failed to delete User")
		return
	}

	response := &util.Response{
		Success: true,
		Message: "Deleted Sucessfully"}
	c.JSON(http.StatusOK, response)
}

// @Summary Update WorkspaceState by it's ID
// @Description Find WorkspaceState by queryID and update it whole
// @Tags WorkspaceState
// @Param id path string true "WorkspaceState ID"
// @Success 200 {object} models.WorkspaceState
// @Failure 400,404 {object} object
// @Router /{id} [put]
func (w *WorkspaceController) UpdateWorkspaceState(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}

	ws := models.Workspace{ID: id}

	wsRecord, err := w.service.Find(&ws)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "User with given id not found")
		return
	}
	ctx.ShouldBindJSON(&wsRecord)

	if ws.Title == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
		return
	}
	if ws.Author.ID == 0 {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Author is required")
		return
	}

	if err := w.service.Update(wsRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store User")
		return
	}
	response := wsRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated User",
		Data:    response,
	})
}
