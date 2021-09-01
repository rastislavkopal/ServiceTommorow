package models

import (
	"time"

	"gorm.io/gorm"
)

// User model definition
type User struct {
	gorm.Model
	ID           uint64    `gorm:"primary_key;auto_increment" json:"id"`
	FullName     string    `gorm:"size:80;not null" json:"fullname" binding:"required"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	Email        string    `gorm:"uniqueIndex:unique_user_email,sort:desc;not null" json:"email" binding:"required,email"`
	PasswordHash string    `gorm:"column:password;not null" json:"password" binding:"required"`
}

// TableName for user table model
func (user *User) TableName() string {
	return "users"
}

// ResponseMap -> response map method of User
func (user *User) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})

	resp["id"] = user.ID
	resp["fullname"] = user.FullName
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
