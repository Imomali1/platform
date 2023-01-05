package handlers

import (
	"github.com/Imomali1/platform/platform/app/handlers/auth"
	"github.com/Imomali1/platform/platform/app/handlers/employees"
)

type Handler struct {
	AuthHandler     *auth.Handler
	EmployeeHandler *employees.Handler
}
