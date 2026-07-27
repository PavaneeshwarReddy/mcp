package app

import (
	"rest-mcp/internal/users"
)

type Dependencies struct {
	UserHandler *users.Handler
}
