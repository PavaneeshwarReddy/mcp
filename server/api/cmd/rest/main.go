package main

import (
	"rest-mcp/internal/router"

	"github.com/bytedance/gopkg/util/logger"
)

/*
Module: Main
Usage: Intilaizes the project
*/

func main() {
	r := router.New() // intialize the global gin engine

	logger.Fatal(
		r.Run(":8080"),
	)
}
