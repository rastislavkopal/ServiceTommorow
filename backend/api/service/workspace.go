package service

import (
	"backend/api/repository"
	"backend/models"
)

type WorkspaceService struct {
	repository *repository.WorkspaceRepository
}

func NewWorkspaceService(r *repository.WorkspaceRepository) WorkspaceService {
	return WorkspaceService{
		repository: r,
	}
}

// Save -> calls workspaceRepository save method
func (w *WorkspaceService) Save(ws models.Workspace) error {
	return w.repository.Save(ws)
}

// FindAll -> calls to workspaceRepository FindAll method
func (w *WorkspaceService) FindAll(ws models.Workspace) (*[]models.Workspace, int64, error) {
	return w.repository.FindAll()
}

// Find -> calls to workspaceRepository Find method
func (w *WorkspaceService) Find(ws *models.Workspace) (*models.Workspace, error) {
	return w.repository.Find(ws)
}

// Update -> calls workspaceRepo update method
func (w *WorkspaceService) Update(ws models.Workspace) error {
	return w.repository.Update(ws)
}

// Delete -> calls workspaceRepository delete method
func (w *WorkspaceService) Delete(id int64) error {
	var workspace models.Workspace
	workspace.ID = id
	return w.repository.Delete(workspace)
}
