package models

import "time"

type User struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	Email        string    `json:"email" gorm:"unique; not null"`
	PasswordHash string    `json:"-" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime;default:CURRENT_TIMESTAMP;not null"`
}
