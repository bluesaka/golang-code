package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	// 无中间件启动
	// r := gin.New()
	// 默认启动方式：包含 Logger、Recovery 中间件
	r := gin.Default()
	r.GET("/get/:id", func(context *gin.Context) {
		// uri参数
		id := context.Param("id")
		// 传参
		uid := context.Query("uid")
		// 默认值
		defaultUid := context.DefaultQuery("uid", "999")
		context.JSON(200, gin.H{
			"id":         id,
			"uid":        uid,
			"defaultUid": defaultUid,
		})
	})

	r.POST("/post", func(context *gin.Context) {
		uid := context.PostForm("uid")
		defaultUid := context.DefaultPostForm("uid", "999")
		context.JSON(200, gin.H{
			"uid":        uid,
			"defaultUid": defaultUid,
		})
	})

	r.DELETE("/delete/:id", func(context *gin.Context) {
		id := context.Param("id")
		context.JSON(200, gin.H{
			"id": id,
		})
	})

	r.PUT("put", func(context *gin.Context) {
		uid := context.PostForm("uid")
		defaultUid := context.DefaultPostForm("uid", "999")
		context.JSON(200, gin.H{
			"uid":        uid,
			"defaultUid": defaultUid,
		})
	})

	r.POST("upload-file", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		_ = context.SaveUploadedFile(file, "123.txt")
		context.JSON(200, gin.H{
			"file": file,
		})
	})

	// 访问静态文件需要先设置路径
	r.Static("/public", "./public")

	middlewareRouter(r)

	// 启动方式1
	//r.Run(":8091")

	// 启动方式2
	//http.ListenAndServe(":8091", r)

	// 启动方式3
	s := http.Server{
		Addr:           ":8091",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func middlewareRouter(r *gin.Engine) {
	r.Use(middleware1, middleware2).Use(middleware3)
	g := r.Group("/group1")
	g.GET("test2", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"test2": "test2",
		})
	})
	{
		g.GET("test1", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"test1": "test1",
			})
		})
	}
}

func middleware1(ctx *gin.Context) {
	log.Println("middleware 1")
	ctx.Next()
}

func middleware2(ctx *gin.Context) {
	log.Println("middleware 2")
	ctx.Next()
}

func middleware3(ctx *gin.Context) {
	log.Println("middleware 3")
	ctx.Next()
}

func cors(ctx *gin.Context) {
	log.Println("middleware cors")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Max-Age", "86400")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
	ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	ctx.Header("Access-Control-Allow-Credentials", "false")
	ctx.Next()
}
