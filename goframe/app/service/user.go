package service

import (
	"github.com/gogf/gf/frame/g"
	"my-goframe/app/dao"
	"my-goframe/app/model"
	"my-goframe/logutil"
)

var User = userService{}

type userService struct{}

func (s *userService) FindByName(name string) (users []model.User, err error) {
	// redis cache
	//_, err = g.Redis().Do("SET", "gf:k1", "v1")
	_, err = g.Redis().DoVar("SET", "gf:k1", "v1")
	if err != nil {
		logutil.Zap.Errorf("[service_user]GetInfo, set redis error: %v", err)
		return
	}

	reply, err := g.Redis().DoVar("GET", "gf:k1")
	if err != nil {
		logutil.Zap.Errorf("[service_user]GetInfo, get redis error: %v", err)
		return
	}
	logutil.Zap.Errorf("[service_user]GetInfo, redis value: %v", reply.String())

	resp, err := dao.User.Where("name = ?", name).All()
	if err != nil {
		logutil.Zap.Errorf("[service_user]GetInfo, get error: %v", err)
		return
	}

	for _, v := range resp {
		users = append(users, model.User{
			Id:       v.Id,
			Name:     v.Name,
			Age:      v.Age,
			Birthday: v.Birthday,
		})
	}
	return
}
