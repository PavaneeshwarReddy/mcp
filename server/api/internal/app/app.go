package app

import (
	"rest-mcp/internal/users"

	"gorm.io/gorm"
)

func Build(db *gorm.DB) *Dependencies {
	userRepo := users.NewRepository(db)
	userSvc := users.NewService(userRepo)
	userHdlr := users.NewHandler(&userSvc)

	return &Dependencies{
		UserHandler: &userHdlr,
	}
}
