package service

import (
	"github.com/kwokky/kais-blog-svc/library/ecode"
	"github.com/kwokky/kais-blog-svc/model"
)

// checkPostExist 检查分类是否存在
func (s *Service) checkCategoryExist(id int64) error {
	var category model.Category
	res := s.Db.First(&category, id)
	if res.RowsAffected <= 0 {
		return ecode.CategoryNotFound
	}
	return nil
}

// checkCategoryName 检查分类名称是否存在
func (s *Service) checkCategoryName(name string, excludeIds []int64) error {
	var (
		category model.Category
		count    int64
	)
	s.Db.Model(&category).Not(excludeIds).Where("name = ?", name).Count(&count)
	if count > 0 {
		return ecode.CategoryExist
	}
	return nil
}
