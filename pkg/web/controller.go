package web

import "github.com/gin-gonic/gin"

type Controller struct {
	router *gin.Engine
}

func NewController(router *gin.Engine) *Controller {
	return &Controller{
		router: router,
	}
}
