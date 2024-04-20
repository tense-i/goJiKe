package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/hello", func(ctx *gin.Context) {
		name := ctx.Query("name")
		ctx.String(http.StatusOK, "name is %s", name)
	})

	server.Run(":8080")
}
