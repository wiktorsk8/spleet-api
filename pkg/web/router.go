package web

import (
	"github.com/gin-gonic/gin"
	"spleet-api/pkg/user"
)

type Router struct{}

func (r *Router) Boot() {

	router := gin.Default()
	api := router.Group("/api")

	userController := user.NewController(router)
	api.GET("/users", userController.GetUser)
}
