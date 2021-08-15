package models

import (
	"time"
)

// User model definition
type User struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string    `gorm:"size:50" json:"first_name"`
	LastName  string    `gorm:"size:50" json:"last_name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password"`
}

// TableName for user table model
func (user *User) TableName() string {
	return "user"
}

// ResponseMap -> response map method of User
func (user *User) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})

	resp["id"] = user.ID
	resp["first_name"] = user.FirstName
	resp["last_name"] = user.LastName
	resp["created_at"] = user.CreatedAt
	resp["updated_at"] = user.UpdatedAt

	return resp
}
