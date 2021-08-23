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
func (w *WorkspaceRepository) Save(ws models.Workspace) error {
	return w.db.DB.Save(&ws).Error
}

// FindAll -> Method for fetching all workspaces
func (w *WorkspaceRepository) FindAll() (*[]models.Workspace, int64, error) {
	var workspaces []models.Workspace

	result := w.db.DB.Find(&workspaces)

	return &workspaces, result.RowsAffected, result.Error
}

// Update -> method for update
func (w *WorkspaceRepository) Update(ws models.Workspace) error {
	return w.db.DB.Save(&ws).Error
}

// Delete -> method to delete user by id
func (w *WorkspaceRepository) Delete(ws models.Workspace) error {
	return w.db.DB.Delete(&ws).Error
}
