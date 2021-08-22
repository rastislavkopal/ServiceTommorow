package models

import "gorm.io/gorm"

// Workspace (board) model definition
type Workspace struct {
	gorm.Model
	ID             int64  `json:"id" gorm:"primary_key;auto_increment"`
	Title          string `json:"title" gorm:"size:50"`
	Description    string `json:"description" gorm:"size:250"`
	OwnerId        User   `json:"owner_id" gorm:"foreignKey:ID"`
	WorkspaceUsers []User `json:"workspace_users,omitempty" gorm:"many2many:workspace_users"`
}

// TableName for Workspace table model
func (ws *Workspace) TableName() string {
	return "workspaces"
}

// ResponseMap -> response map method of workspace
func (ws *Workspace) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})

	resp["id"] = ws.ID
	resp["title"] = ws.Title
	resp["description"] = ws.Description
	resp["owner_id"] = ws.OwnerId.ID
	// resp["workspace_users"] = ws.

	return resp
}

// WorkspaceState model definition e.g. TODO / In-progress / done etc.
type WorkspaceState struct {
	gorm.Model
	ID          uint64    `json:"id"`
	WorkspaceId Workspace `json:"workspace_id" gorm:"foreignKey:ID"`
	StateTitle  string    `json:"state_title"`
	StateOrder  int       `json:"state_order"`
}

// TableName for Workspace table model
func (wss *WorkspaceState) TableName() string {
	return "workspace_states"
}

// ResponseMap -> response map method of workspace
func (wss *WorkspaceState) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})

	resp["id"] = wss.ID
	resp["workspace_id"] = wss.WorkspaceId.ID
	resp["state_title"] = wss.StateTitle
	resp["state_order"] = wss.StateOrder

	return resp
}
