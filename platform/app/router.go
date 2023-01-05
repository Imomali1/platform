package app

import (
	"github.com/Imomali1/platform/platform/app/handlers"
	"github.com/Imomali1/platform/platform/app/handlers/auth"
	"github.com/Imomali1/platform/platform/app/handlers/employees"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	var h = handlers.Handler{
		AuthHandler:     auth.NewHandler(),
		EmployeeHandler: employees.NewHandler(),
	}

	// no auth required
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to Platform!"})
	})

	router.POST("/sign-in", h.AuthHandler.SignIn)
	router.POST("/sign-up", h.AuthHandler.SignUp)

	api := router.Group("/api/v1")
	{
		employee := api.Group("/employees")
		{
			employee.POST("/create", h.EmployeeHandler.Create)
			employee.GET("/list", h.EmployeeHandler.List)
			employee.GET("/detail/:id", h.EmployeeHandler.Detail)
			employee.PUT("/update/:id", h.EmployeeHandler.Update)
			employee.DELETE("/delete/:id", h.EmployeeHandler.Delete)
		}
	}

	return router
}
