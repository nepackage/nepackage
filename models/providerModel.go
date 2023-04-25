package models

import "time"

type Provider struct {
	ID        uint      `json:"id" gorm:"primary_key; not null; autoIncrement"`
	Name      string    `json:"name" gorm:"not null"`
	Type      string    `json:"type" gorm:"not null"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"not null"`
}
