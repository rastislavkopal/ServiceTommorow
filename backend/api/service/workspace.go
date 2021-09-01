package service

import (
	"backend/api/repository"
	"backend/models"
	"errors"
	"strconv"
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
func (w *WorkspaceService) Save(ws *models.Workspace, user_id uint64) error {
	user := models.User{ID: user_id}

	foundUser, err := w.userRepo.Find(user)
	if err != nil {
		return errors.New("could not find User ID: " + strconv.Itoa(int(user_id)))
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
func (w *WorkspaceService) Update(ws *models.Workspace) error {
	return w.repository.Update(ws)
}

// Delete -> calls workspaceRepository delete method
func (w *WorkspaceService) Delete(id uint64) error {
	var workspace models.Workspace
	workspace.ID = id
	return w.repository.Delete(workspace)
}

//****************************
//*  WorkspaceState handlers *
//****************************

// Save -> calls WorkspaceState save method
func (w *WorkspaceService) SaveWorkspaceState(wss *models.WorkspaceState, workspace_id uint64) error {
	var ws models.Workspace

	foundWorkspace, err := w.repository.Find(&ws)
	if err != nil {
		return errors.New("could not find workspaceState")
	}

	wss.Workspace = *foundWorkspace

	return w.repository.SaveWorkspaceState(wss)
}

// FindAll -> calls to workspaceState FindAll method
func (w *WorkspaceService) FindAllWorkspaceState(ws models.WorkspaceState) (*[]models.WorkspaceState, int64, error) {
	return w.repository.FindAllWorkspaceState()
}

// Find -> calls to workspaceState Find method
func (w *WorkspaceService) FindWorkspaceState(ws *models.WorkspaceState) (*models.WorkspaceState, error) {
	return w.repository.FindWorkspaceState(ws)
}

// Update -> calls workspaceState update method
func (w *WorkspaceService) UpdateWorkspaceState(ws *models.WorkspaceState) error {
	return w.repository.UpdateWorkspaceState(ws)
}

// Delete -> calls workspaceState delete method
func (w *WorkspaceService) DeleteWorkspaceState(id uint64) error {
	var workspaceState models.WorkspaceState
	workspaceState.ID = id
	return w.repository.DeleteWorkspaceState(&workspaceState)
}
