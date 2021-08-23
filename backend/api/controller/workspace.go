package controller

import (
	"backend/api/service"
	"backend/models"
	"backend/util"
	"net/http"

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
// @Success 200 {array} models.User
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
		Message: "User result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}
