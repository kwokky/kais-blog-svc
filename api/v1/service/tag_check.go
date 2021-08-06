package service

import (
	"github.com/kwokky/kais-blog-svc/library/ecode"
	"github.com/kwokky/kais-blog-svc/model"
)

// checkTagExist 检查标签是否存在
func (s *Service) checkTagExist(id int64) error {
	var tag model.Tag
	res := s.Db.First(&tag, id)
	if res.RowsAffected <= 0 {
		return ecode.TagNotFound
	}
	return nil
}

// checkTagName 检查标签名称是否存在
func (s *Service) checkTagName(name string, excludeIds []int64) error {
	var (
		tag model.Tag
		count    int64
	)
	s.Db.Model(&tag).Not(excludeIds).Where("name = ?", name).Count(&count)
	if count > 0 {
		return ecode.TagExist
	}
	return nil
}
