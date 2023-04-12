package models

import "time"

type GithubCredential struct {
	ID        uint      `json:"id" gorm:"primary_key; not null; autoIncrement"`
	Name      string    `json:"name" gorm:"not null"`
	Username  string    `json:"username" gorm:"not null"`
	Password  string    `json:"-" gorm:"not null"`
	State     bool      `json:"state" gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

type GithubCredentialCreate struct {
	ID        uint   `json:"id"`
	Name      string `json:"name" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	State     bool   `json:"state"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GithubCredentialUpdate struct {
	ID        uint   `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	State     bool   `json:"state,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
