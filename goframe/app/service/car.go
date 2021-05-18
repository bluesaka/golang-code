package service

import (
	"github.com/gogf/gf/frame/g"
	"log"
	"my-goframe/app/model"
	"my-goframe/logutil"
)

var Car carService

type carService struct {
	name  string
	price int
}

func init() {
	log.Println("[service]test init")
	logutil.Zap.Info("[service] test init with zap")

	// read config
	name := g.Cfg("car").GetString("porsche.name")
	price := g.Cfg("car").GetInt("porsche.price")
	logutil.Zap.Infof("porsche model: %s, price: %d", name, price)

	// dump config
	g.Cfg("car").Dump()

	// init Car
	Car = carService{
		name:  name,
		price: price,
	}
}

func (s *carService) GetCarInfo(req *model.CarInfoReq) (ret model.CarInfoResp, err error) {
	ret = model.CarInfoResp{
		ID:    req.ID,
		Name:  s.name,
		Price: s.price,
	}
	return
}
