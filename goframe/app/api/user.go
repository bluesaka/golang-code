package api

import (
	"github.com/gogf/gf/net/ghttp"
	"my-goframe/app/service"
	"my-goframe/library/response"
)

var User = new(userApi)

type userApi struct{}

func (a *userApi) FindByName(r *ghttp.Request) {
	name := r.GetString("name", "")
	if name == "" {
		response.JsonExit(r, 1001, "name empty")
	}

	resp, err := service.User.FindByName(name)
	if err != nil {
		response.JsonExit(r, 1002, err.Error())
	}

	response.Json(r, 0, "success", resp)
}
