package service

import (
	"backend/api/repository"
	"backend/models"
	"errors"
)

type WorkspaceService struct {
	repository *repository.WorkspaceRepository
	userRepo   *repository.UserRepository
}

func NewWorkspaceService(r *repository.WorkspaceRepository, u *repository.UserRepository) WorkspaceService {
	return WorkspaceService{
		repository: r,
		userRepo:   u,
	}
}

// Save -> calls workspaceRepository save method
func (w *WorkspaceService) Save(ws models.Workspace, user_id uint64) error {
	var user models.User
	user.ID = user_id

	foundUser, err := w.userRepo.Find(user)
	if err != nil {
		return errors.New("could not find user")
	}

	ws.Author = foundUser

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
func (w *WorkspaceService) Delete(id uint64) error {
	var workspace models.Workspace
	workspace.ID = id
	return w.repository.Delete(workspace)
}
