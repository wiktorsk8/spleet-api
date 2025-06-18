package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"spleet-api/pkg/web"
)

func main() {
	router := web.Router{}

	router.Boot()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
