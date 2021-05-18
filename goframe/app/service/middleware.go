package service

import (
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

var Middleware = middlewareService{}

type middlewareService struct{}

func (s *middlewareService) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func (s *middlewareService) MyCORS(r *ghttp.Request) {
	o := r.Response.DefaultCORSOptions()
	o.AllowOrigin = "*"
	o.AllowHeaders = "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token"
	o.MaxAge = 86400
	o.AllowMethods = "GET, POST, OPTIONS, PUT, PATCH, DELETE"
	o.ExposeHeaders = "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type"
	o.AllowCredentials = "false"

	r.Middleware.Next()
}

func (s *middlewareService) Auth(r *ghttp.Request) {
	if true {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatus(http.StatusForbidden)
	}
}
