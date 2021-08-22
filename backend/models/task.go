package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID               int64          `json:"id"`
	Title            string         `json:"title"`
	Description      string         `json:"description"`
	Status           string         `json:"status"`
	OwnerId          User           `json:"owner_id" gorm:"foreignKey:ID"`
	WorkspaceId      Workspace      `json:"workspace_id" gorm:"foreignKey:ID"`
	Deadline         time.Time      `json:"deadline"`
	WorkspaceStateId WorkspaceState `json:"workspace_state_id" gorm:"foreignKey:ID"`
	Users            []User         `json:"users,omitempty" gorm:"many2many:tasks_users"`
}

// TableName for Tasks table model
func (task *Task) TableName() string {
	return "tasks"
}

// ResponseMap -> response map method of tasks
func (task *Task) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})

	resp["id"] = task.ID
	resp["title"] = task.Title
	resp["description"] = task.Description
	resp["status"] = task.Status
	resp["owner_id"] = task.OwnerId.ID
	resp["workspace_id"] = task.WorkspaceId.ID
	resp["deadline"] = task.Deadline
	resp["workspace_state_id"] = task.WorkspaceStateId.ID

	return resp
}
