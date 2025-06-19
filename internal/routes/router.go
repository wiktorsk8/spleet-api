package routes

import (
	"github.com/gin-gonic/gin"
	"spleet-api/internal/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/ping", handlers.GetUser)
	}

	return r
}
