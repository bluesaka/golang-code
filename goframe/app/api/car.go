package api

import (
	"github.com/gogf/gf/net/ghttp"
	"my-goframe/app/model"
	"my-goframe/app/service"
	"my-goframe/library/response"
	"my-goframe/logutil"
)

var Car carApi

type carApi struct{}

func (a *carApi) GetCarInfo(r *ghttp.Request) {
	// r.GetXXX
	//id := r.GetInt("id")
	//notExistString := r.GetString("not_exist", "default_value")
	//logutil.Zap.Infof("[api_car]GetCarInfo, id: %d", id)
	//logutil.Zap.Infof("[api_car]GetCarInfo, notExistString: %s", notExistString)

	// r.Parse
	var req *model.CarInfoReq
	if err := r.Parse(&req); err != nil {
		logutil.Zap.Errorf("[api_car]GetCarInfo, Parse error: %v", err)
		return
	}
	logutil.Zap.Errorf("[api_car]GetCarInfo, Parse data: %+v", req)

	resp, err := service.Car.GetCarInfo(req)
	if err != nil {
		response.JsonExit(r, 1001, err.Error())
	}

	response.Json(r, 0, "success", resp)
}
