// Package routes contains every available REST API route
package routes

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// InitRouter will return HTTP gin router
func InitRouter() *gin.Engine {
	router := gin.Default()
	//router.Static("/app/", "./web/app/dist")
	router.Use(static.Serve("/", static.LocalFile("./web/app/dist", false)))
	router.NoRoute(func(c *gin.Context) {
		c.File("./web/app/dist/index.html")
	})
	v1 := router.Group("/api/v1")
	initAuthRoutes(v1)

	return router
}
