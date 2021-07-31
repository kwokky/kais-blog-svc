package service

import (
	"github.com/kwokky/kais-blog-svc/model"
	"gorm.io/gorm"
)

type Service struct {
	Db *gorm.DB
}

func New() *Service {
	srv := &Service{
		Db: model.DB(),
	}
	return srv
}
