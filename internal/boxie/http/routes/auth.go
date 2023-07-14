// Package routes contains every available REST API route
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/twelvee/boxie/internal/boxie/http/controllers"
	"github.com/twelvee/boxie/internal/boxie/http/middlewares"
)

func initAuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/user", controllers.Login)
	rg.DELETE("/user", controllers.DeleteUser)
	rg.POST("/user/invite/accept", controllers.AcceptInvite)
	rg.PUT("/user", controllers.SetUserPassword)

	// Auth area
	rg.GET("/user", middlewares.HasToken(), controllers.GetUser)
	rg.POST("/user/create", middlewares.HasToken(), controllers.Register)
}
