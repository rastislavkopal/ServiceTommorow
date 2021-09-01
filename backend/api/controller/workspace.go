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

	data, total, err := w.service.FindAll(workspaces, ctx)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find workspace: "+err.Error())
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

// CreateWorkspace : CreateWorkspace controller
// @Summary Create new workspace
// @Description Create new workspace
// @Tags Workspaces
// @Success 201 {array} models.Workspace
// @Failure 400 {object} object
// @Router / [post]
func (w *WorkspaceController) CreateWorkspace(ctx *gin.Context) {
	var ws models.Workspace
	ctx.ShouldBindJSON(&ws)

	if ws.Title == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
		return
	}

	// user_id, err := strconv.ParseUint(ctx.PostForm("user_id"), 10, 64)
	if ws.Author.ID == 0 {
		util.ErrorJSON(ctx, http.StatusBadRequest, "user_id is required")
		return
	}

	err := w.service.Save(&ws, ws.Author.ID)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create workspace: "+err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg":          "Successfully created new workspace",
		"workspace_id": ws.ID,
		"author_id":    ws.AuthorID,
		"title":        ws.Title,
		"description":  ws.Description,
	})
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
// @Description Find workspace by queryID and delete it with all it's tasks
// @Tags Workspaces
// @Param id path string true "Workspace ID"
// @Success 200 {object} models.Workspace
// @Failure 400,404 {object} object
// @Router /{id} [delete]
func (w *WorkspaceController) DeleteWorkspace(c *gin.Context) {
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

// @Summary Update workspace by it's ID
// @Description Find workspace by queryID and update it whole
// @Tags Workspaces
// @Param id path string true "Workspace ID"
// @Success 200 {object} models.Workspace
// @Failure 400,404 {object} object
// @Router /{id} [put]
func (w *WorkspaceController) UpdateWorkspace(ctx *gin.Context) {
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
