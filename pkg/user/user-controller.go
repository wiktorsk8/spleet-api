package user

import (
	"net/http"
	"spleet-api/pkg/web"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	web.Controller
}

func (uc *UserController) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id": "uuid-123",
	})
}
