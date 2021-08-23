package controller

import (
	"backend/api/service"
	"backend/models"
	"backend/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WorkspaceController struct {
	service *service.WorkspaceService
}

func NewWorkspaceController(s *service.WorkspaceService) WorkspaceController {
	return WorkspaceController{
		service: s,
	}
}

// GetWorkspaces : GetWorkspaces controller
// @Summary Get list of workspaces
// @Description get all workspaces
// @Tags Workspaces
// @Success 200 {array} models.Workspace
// @Failure 404 {object} object
// @Router / [get]
func (w *WorkspaceController) GetWorkspaces(ctx *gin.Context) {
	var workspaces models.Workspace

	// keyword := ctx.Query("keyword")

	data, total, err := w.service.FindAll(workspaces)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find questions")
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

// AddWorkspace : AddWorkspace controller
// @Summary Create new workspace
// @Description Create new workspace
// @Tags Workspaces
// @Success 201 {array} models.Workspace
// @Failure 400 {object} object
// @Router / [post]
func (w *WorkspaceController) AddWorkspace(ctx *gin.Context) {
	var ws models.Workspace
	ctx.ShouldBindJSON(&ws)

	if ws.Title == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
		return
	}

	err := w.service.Save(ws)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create workspace")
		return
	}
	util.SuccessJSON(ctx, http.StatusCreated, "Successfully created new workspace")
}

// GetWorkspace -> get workspace by id
// @Summary Get one workspace
// @Description get workspace by ID
// @Tags Workspaces
// @Param id path string true "Workspace ID"
// @Success 200 {object} models.Workspace
// @Failure 400,404 {object} object
// @Router /{id} [get]
func (w *WorkspaceController) GetWorkspace(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	var workspace models.Workspace
	workspace.ID = id
	foundWorkspace, err := w.service.Find(&workspace)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Error Finding Workspace")
		return
	}
	response := foundWorkspace.ResponseMap()

	c.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of Workspace",
		Data:    &response})

}
