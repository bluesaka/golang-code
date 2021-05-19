package service

import "my-gin/dao"

type Service struct {
	dao *dao.Dao
}

func Init() *Service {
	return &Service{
		dao: dao.NewDao(),
	}
}
