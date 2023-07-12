// Package routes contains every available REST API route
package routes

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// InitRouter will return HTTP gin router
func InitRouter(webAppPath string) *gin.Engine {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile(webAppPath, false)))
	router.NoRoute(func(c *gin.Context) {
		c.File(webAppPath + "/index.html")
	})
	v1 := router.Group("/api/v1")
	initSetupRoutes(v1)
	initAuthRoutes(v1)
	initUsersRoutes(v1)
	initClustersRoutes(v1)
	initBoxesRoutes(v1)
	initEnvironmentsRoutes(v1)

	return router
}
