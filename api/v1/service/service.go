package service

import (
	"github.com/kwokky/kais-blog-svc/model"
	"gorm.io/gorm"
)

type Service struct {
	Db *gorm.DB
}

func New() *Service {
	svc := &Service{
		Db: model.DB(),
	}
	return svc
}
