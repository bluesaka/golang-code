package routers

import (
	"github.com/gin-gonic/gin"
	"my-gin/web/controllers"
)

func SetRouters(g *gin.Engine) {
	//g.Static("/public", "public")

	api := g.Group("/api")
	{
		api.GET("/login", controllers.UserLogin)
	}
}
