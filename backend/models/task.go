package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID               uint64         `json:"id"`
	Title            string         `json:"title"`
	Description      string         `json:"description"`
	Status           string         `json:"status"`
	Author           User           `json:"author_id" gorm:"foreignKey:AuthorID"`
	AuthorID         uint64         `json:"-"`
	Workspace        Workspace      `json:"workspace_id" gorm:"foreignKey:WorkspaceID"`
	WorkspaceID      uint64         `json:"-"`
	Deadline         time.Time      `json:"deadline"`
	WorkspaceState   WorkspaceState `json:"workspace_state_id" gorm:"foreignKey:ID"`
	WorkspaceStateID uint64         `json:"-"`
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
	resp["author_id"] = task.Author.ID
	resp["workspace_id"] = task.Workspace.ID
	resp["deadline"] = task.Deadline
	resp["workspace_state_id"] = task.WorkspaceState.ID

	return resp
}
