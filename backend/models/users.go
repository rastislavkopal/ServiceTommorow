package models

import (
	"time"

	"gorm.io/gorm"
)

// User model definition
type User struct {
	gorm.Model
	ID           int64     `gorm:"primary_key;auto_increment" json:"id"`
	FirstName    string    `gorm:"size:50" json:"first_name"`
	LastName     string    `gorm:"size:50" json:"last_name"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	Email        string    `gorm:"unique_index;not null" json:"email"`
	PasswordHash string    `gorm:"column:password;not null" json:"password"`
}

// TableName for user table model
func (user *User) TableName() string {
	return "users"
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

// TokenDetails to define AccessToken token and Refresh token
type TokenDetails struct {
	gorm.Model
	AccessToken  string `gorm:"size:500" json:"access_token"`
	RefreshToken string `gorm:"size:500" json:"refresh_token"`
	AccessUuid   string `gorm:"size:40" json:"access_uuid"`
	RefreshUuid  string `gorm:"size:40" json:"refresh_uuid"`
	AtExpires    int64  `json:"at_expires"`
	RtExpires    int64  `json:"rt_expires"`
}

// TableName for TokenDetails table model
func (tokenDetails *TokenDetails) TableName() string {
	return "tokens"
}

// ResponseMap -> response map method of TokenDetails
func (tokenDetails *TokenDetails) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})

	resp["access_token"] = tokenDetails.AccessToken
	resp["refresh_token"] = tokenDetails.RefreshToken
	resp["access_uuid"] = tokenDetails.AccessUuid
	resp["refresh_uuid"] = tokenDetails.RefreshUuid
	resp["at_expires"] = tokenDetails.AtExpires
	resp["rt_expires"] = tokenDetails.RtExpires

	return resp
}
