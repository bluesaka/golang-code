package main

import (
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	_ "my-gin/demo/boot"
	"my-gin/demo/utils"
	"my-gin/demo/web/routers"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	routers.SetRouters(r)

	if !viper.GetBool("app.debug") {
		gin.SetMode(gin.ReleaseMode)
	}

	addr := viper.GetString("app.addr")
	srv := &http.Server{
		Addr:           addr,
		Handler:        r,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	utils.Log.Info("server start on", addr)

	if err := gracehttp.Serve(srv); err != nil {
		panic(fmt.Sprintf("listen error: %s\n", err))
	}
}
