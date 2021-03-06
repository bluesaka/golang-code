package routers

import (
	"github.com/gin-gonic/gin"
	"my-gin/demo/web/controllers"
)

func SetRouters(g *gin.Engine) {
	//g.Static("/public", "public")
	//g.LoadHTMLGlob("public/templates/*.html")

	api := g.Group("/api")
	{
		api.GET("/user/login", controllers.UserLogin)
		api.GET("/user/info", controllers.UserInfo)

		api.GET("/user/http-get", controllers.HttpGet)
		api.GET("/user/http-post", controllers.HttpPost)
	}
}
