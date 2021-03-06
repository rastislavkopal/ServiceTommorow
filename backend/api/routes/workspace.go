package routes

import (
	"backend/api/controller"
	"backend/common"
)

type WorkspaceRoute struct {
	controller *controller.WorkspaceController
	handler    *common.GinRouter
}

func NewWorkspaceRoute(c *controller.WorkspaceController, h *common.GinRouter) WorkspaceRoute {
	return WorkspaceRoute{
		controller: c,
		handler:    h,
	}
}

func (w *WorkspaceRoute) Setup() {
	workspace := w.handler.Gin.Group("/workspace")
	{
		workspace.GET("/", w.controller.GetWorkspaces)
		workspace.POST("/", w.controller.CreateWorkspace)
		workspace.GET("/:id", w.controller.GetWorkspace)
		workspace.DELETE("/:id", w.controller.DeleteWorkspace)
		workspace.PUT("/:id", w.controller.UpdateWorkspace)
	}

	workspaceState := w.handler.Gin.Group("/workspace/:id/state")
	{
		workspaceState.GET("/", w.controller.GetWorkspaceStates)
		workspaceState.POST("/", w.controller.CreateWorkspaceState)
		workspaceState.GET("/:wss_id", w.controller.GetWorkspaceState)
		workspaceState.DELETE("/:wss_id", w.controller.DeleteWorkspaceState)
		workspaceState.PUT("/:wss_id", w.controller.UpdateWorkspaceState)
	}
}
