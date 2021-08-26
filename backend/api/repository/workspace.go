package repository

import (
	"backend/common"
	"backend/models"
)

type WorkspaceRepository struct {
	db *common.Database
}

func NewWorkspaceRepository(db *common.Database) WorkspaceRepository {
	return WorkspaceRepository{
		db: db,
	}
}

// Save -> function to save workspace into DB
func (w *WorkspaceRepository) Save(ws *models.Workspace) error {
	return w.db.DB.Create(ws).Error
}

// FindAll -> Method for fetching all workspaces
func (w *WorkspaceRepository) FindAll() (*[]models.Workspace, int64, error) {
	var workspaces []models.Workspace

	result := w.db.DB.Find(&workspaces)

	return &workspaces, result.RowsAffected, result.Error
}

// Find-> Method for fetching workspace by id
func (w *WorkspaceRepository) Find(ws *models.Workspace) (*models.Workspace, error) {
	var foundWs models.Workspace

	err := w.db.DB.First(&foundWs, "id = ?", ws.ID).Error

	return &foundWs, err
}

// Update -> method for update
func (w *WorkspaceRepository) Update(ws *models.Workspace) error {
	return w.db.DB.Save(ws).Error
}

// Delete -> method to delete workspace by id
func (w *WorkspaceRepository) Delete(ws models.Workspace) error {
	return w.db.DB.Delete(&ws).Error
}

//****************************
//*  WorkspaceState handlers *
//****************************
// Save -> function to save workspace into DB
func (w *WorkspaceRepository) SaveWorkspaceState(wss *models.WorkspaceState) error {
	return w.db.DB.Create(wss).Error
}

// FindAll -> Method for fetching all WorkspacesStates
func (w *WorkspaceRepository) FindAllWorkspaceState() (*[]models.WorkspaceState, int64, error) {
	var workspacesStates []models.WorkspaceState

	result := w.db.DB.Find(&workspacesStates)

	return &workspacesStates, result.RowsAffected, result.Error
}

// Find-> Method for fetching WorkspacesStates by id
func (w *WorkspaceRepository) FindWorkspaceState(wss *models.WorkspaceState) (*models.WorkspaceState, error) {
	var foundWss models.WorkspaceState

	err := w.db.DB.First(&foundWss, "id = ?", wss.ID).Error

	return &foundWss, err
}

// Update -> method for update WorkspacesStates
func (w *WorkspaceRepository) UpdateWorkspaceState(wss *models.WorkspaceState) error {
	return w.db.DB.Save(wss).Error
}

// Delete -> method to delete WorkspacesStates by id
func (w *WorkspaceRepository) DeleteWorkspaceState(wss *models.WorkspaceState) error {
	return w.db.DB.Delete(wss).Error
}
