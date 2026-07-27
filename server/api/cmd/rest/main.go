package main

import (
	"rest-mcp/internal/app"
	"rest-mcp/internal/database"
	"rest-mcp/internal/router"
	"rest-mcp/internal/users"

	"github.com/bytedance/gopkg/util/logger"
)

/*
Module: Main
Usage: Intilaizes the project
*/

func main() {

	cfg := database.LoadConfig()         // load database config
	db, err := database.NewPostgres(cfg) // intialize gorm database
	logger.Fatal(err)

	logger.Fatal(db.AutoMigrate( // handled automigration of models
		&users.User{},
	))

	app := app.Build(db) // add all dependencies

	r := router.New(db, app) // intialize the global gin engine

	logger.Fatal(r.Run(":8080"))
}
