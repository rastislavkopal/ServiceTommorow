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
// @Router /{id} [get]
func (w *WorkspaceController) GetWorkspaceStates(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}

	ws := models.Workspace{ID: id}
	foundWs, err := w.service.Find(&ws)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusNotFound, "Failed to find Workspace")
		return
	}

	var wss models.WorkspaceState
	wss.Workspace = *foundWs

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
// @Router /{id} [post]
func (w *WorkspaceController) CreateWorkspaceState(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}

	ws := models.Workspace{ID: id}
	foundWs, err := w.service.Find(&ws)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusNotFound, "Failed to find Workspace")
		return
	}

	var wss models.WorkspaceState
	ctx.ShouldBindJSON(&ws)
	if wss.Title == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
		return
	}

	err = w.service.SaveWorkspaceState(&wss, foundWs.ID)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create WorkspaceState")
		return
	}

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully created new WorkspaceState",
		Data: map[string]interface{}{
			"workspace_state_id": wss.ID,
		}})
}

// @Summary Get one WorkspaceState
// @Description get WorkspaceState by ID
// @Tags WorkspaceState
// @Param id path string true "WorkspaceState ID"
// @Success 200 {object} models.WorkspaceState
// @Failure 400,404 {object} object
// @Router /{id}/state/{wss_id} [get]
func (w *WorkspaceController) GetWorkspaceState(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64) //type conversion string to uint64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}

	wssIdParam := ctx.Param("wss_id")
	wssId, err := strconv.ParseUint(wssIdParam, 10, 64) //type conversion string to uint64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}

	ws := models.Workspace{ID: id}
	foundWs, err := w.service.Find(&ws)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusNotFound, "Failed to find Workspace")
		return
	}

	wss := models.WorkspaceState{ID: wssId, WorkspaceID: foundWs.ID}
	_, err = w.service.FindWorkspaceState(&wss)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create WorkspaceState")
		return
	}

	response := wss.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result WorkspaceState",
		Data:    &response})
}

// @Summary Delete workspace by it's ID
// @Description Find WorkspaceState by queryID and delete it
// @Tags WorkspaceState
// @Param id path string true "WorkspaceState ID"
// @Success 200 {object} models.WorkspaceState
// @Failure 400,404 {object} object
// @Router /{id}/state/{wss_id} [delete]
func (w *WorkspaceController) DeleteWorkspaceState(c *gin.Context) {
	// idParam := c.Param("id")
	// id, err := strconv.ParseUint(idParam, 10, 64) //type conversion string to uint64
	// if err != nil {
	// 	util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
	// 	return
	// }

	// err = w.service.Delete(id)

	// if err != nil {
	// 	util.ErrorJSON(c, http.StatusNotFound, "Failed to delete User")
	// 	return
	// }

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
// @Router /{id}/state/{wss_id} [put]
func (w *WorkspaceController) UpdateWorkspaceState(ctx *gin.Context) {
	// idParam := ctx.Param("id")

	// id, err := strconv.ParseUint(idParam, 10, 64)

	// if err != nil {
	// 	util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
	// 	return
	// }

	// ws := models.Workspace{ID: id}

	// wsRecord, err := w.service.Find(&ws)

	// if err != nil {
	// 	util.ErrorJSON(ctx, http.StatusBadRequest, "User with given id not found")
	// 	return
	// }
	// ctx.ShouldBindJSON(&wsRecord)

	// if ws.Title == "" {
	// 	util.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
	// 	return
	// }
	// if ws.Author.ID == 0 {
	// 	util.ErrorJSON(ctx, http.StatusBadRequest, "Author is required")
	// 	return
	// }

	// if err := w.service.Update(wsRecord); err != nil {
	// 	util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store User")
	// 	return
	// }
	// response := wsRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated User",
		// Data:    response,
	})
}
