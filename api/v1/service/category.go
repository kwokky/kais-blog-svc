package service

import (
	rspModel "github.com/kwokky/kais-blog-svc/api/v1/model"
	"github.com/kwokky/kais-blog-svc/api/v1/param"
	"github.com/kwokky/kais-blog-svc/library/ecode"
	"github.com/kwokky/kais-blog-svc/model"
	"log"
)

func (s *Service) CategoryCreate(params param.CategoryCreateParams) (err error) {
	var category model.Category

	if params.Name == "" {
		err = ecode.CategoryNameEmpty
		return
	}

	// 分类不能重名
	if err = s.checkCategoryName(params.Name, nil); err != nil {
		return
	}

	category.Name = params.Name
	res := s.Db.Create(&category)
	if res.Error != nil {
		log.Printf("新增分类失败 %s", err)
		return ecode.CategoryCreateError
	}

	return
}

func (s *Service) CategoryUpdate(id int64, params param.CategoryUpdateParams) (err error) {
	if id == 0 {
		err = ecode.CategoryParamError
		return
	}

	var category model.Category
	res := s.Db.First(&category, id)
	if res.RowsAffected <= 0 {
		err = ecode.CategoryNotFound
		return
	}

	// 分类不能重名
	if err = s.checkCategoryName(params.Name, []int64{id}); err != nil {
		return
	}

	updated := s.Db.Model(&category).Updates(model.Category{Name: params.Name})
	if updated.RowsAffected <= 0 {
		log.Printf("修改分类失败 %s", updated.Error)
		err = ecode.CategoryUpdateError
		return
	}

	return
}

func (s *Service) CategoryList(params param.CategoryListParams) (*rspModel.CategoryList, error) {
	var (
		cmodel     model.Category
		categories []*model.Category
		total      int64
		field      = []string{"id", "name"}
		offset     = (params.Page - 1) * params.Size
	)

	res := s.Db.Model(&cmodel).Offset(int(offset)).Limit(int(params.Size)).Select(field).Find(&categories)
	if res.RowsAffected <= 0 {
		return rspModel.NewCategoryListResponse([]*model.Category{}, 0), nil
	}

	s.Db.Model(&cmodel).Count(&total)

	return rspModel.NewCategoryListResponse(categories, total), nil
}

func (s *Service) CategoryDetail(params param.CategoryDetailParams) (detail *rspModel.Category, err error) {
	if params.ID == 0 {
		err = ecode.CategoryParamError
		return
	}

	var category model.Category
	res := s.Db.First(&category, params.ID)
	if res.RowsAffected <= 0 {
		err = ecode.CategoryNotFound
		return
	}

	detail = rspModel.NewCategoryDetailResponse(&category)
	return
}

func (s *Service) CategoryDelete(params param.CategoryDeleteParams) (err error) {
	if params.ID == 0 {
		err = ecode.CategoryParamError
		return
	}

	var category model.Category
	res := s.Db.Delete(&category, params.ID)
	if res.RowsAffected <= 0 {
		err = ecode.CategoryDeleteError
		return
	}

	return
}
