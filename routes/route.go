package routes

import (
	"github.com/baguseka01/golang-jwt-authentication/auth"
	"github.com/baguseka01/golang-jwt-authentication/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/login", controllers.Login)
		api.POST("/register", controllers.Register)
		secured := api.Group("/secured").Use(auth.Authenticate())
		{
			secured.GET("/home", controllers.Home)
		}
	}
	return router
}
