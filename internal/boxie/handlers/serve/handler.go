// Package serve is used to process serve command
package serve

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/twelvee/boxie/internal/boxie/http/routes"
	"net/http"
)

// HandleServeCommand is the boxie serve command handler
func HandleServeCommand(context context.Context, host string, port string, static string) {
	engine := routes.InitRouter(static)
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	if err := engine.Run(host + ":" + port); err != nil {
		panic(err)
	}
}
