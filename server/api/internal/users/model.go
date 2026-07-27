package users

import "gorm.io/gorm"

/*
Module: GORM models for authentication
Usage: Exact replica of database tables
*/

type User struct {
	gorm.Model

	Username string  `gorm:"index; unique"`
	Password string  `gorm:"not null"`
	Email    *string `gorm:"unique"`
	Age      uint    `gorm:"not null"`
}
