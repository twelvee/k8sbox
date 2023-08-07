// Package routes contains every available REST API route
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/twelvee/boxie/internal/boxie/http/controllers"
	"github.com/twelvee/boxie/internal/boxie/http/middlewares"
)

func initUsersRoutes(rg *gin.RouterGroup) {
	// Auth area
	rg.GET("/users", middlewares.HasToken(), controllers.GetUsers)
}
