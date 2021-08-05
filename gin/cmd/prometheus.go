package main

import (
	"github.com/gin-gonic/gin"
	"my-gin/web/middleware"
)

func main() {
	r := gin.New()

	// Prometheus middleware 1
	p := middleware.NewPrometheus("gin")
	p.Use(r)

	// Prometheus middleware 2
	r.Use(middleware.Prometheus2)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "hello world")
	})
	r.Run(":30000")
}
