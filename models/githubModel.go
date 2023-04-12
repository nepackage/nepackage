package models

import "time"

type GithubCredential struct {
	ID        uint      `json:"id" gorm:"primary_key; not null; autoIncrement"`
	Name      string    `json:"name" binding:"required" gorm:"not null"`
	Username  string    `json:"username" binding:"required" gorm:"not null"`
	Password  string    `json:"-" binding:"required" gorm:"not null"`
	State     bool      `json:"state" binding:"required" gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

type GithubCredentialCreate struct {
	ID        uint   `json:"id"`
	Name      string `json:"name" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	State     bool   `json:"state" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GithubCredentialUpdate struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	State     bool   `json:"state"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
