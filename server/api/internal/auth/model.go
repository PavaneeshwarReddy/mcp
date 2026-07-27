package auth

import (
	"time"

	"gorm.io/gorm"
)

/*
Module: GORM models for authentication
Usage: Exact replica of database tables
*/

type RefreshToken struct {
	gorm.Model

	UserId    uint      `gorm:"not null; index"`
	TokenHash string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
	RevokedAt *string
	IPAddress string
}
