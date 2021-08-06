package service

import (
	rspModel "github.com/kwokky/kais-blog-svc/api/v1/model"
	"github.com/kwokky/kais-blog-svc/api/v1/param"
	"github.com/kwokky/kais-blog-svc/library/ecode"
	"github.com/kwokky/kais-blog-svc/model"
	"log"
)

func (s *Service) TagCreate(params param.TagCreateParams) (err error) {
	if params.Name == "" {
		err = ecode.TagNameEmpty
		return
	}

	// 标签不能重名
	if err = s.checkTagName(params.Name, nil); err != nil {
		return
	}

	tag := model.Tag{Name: params.Name}
	res := s.Db.Create(&tag)
	if res.Error != nil {
		log.Printf("新增标签失败 %s", err)
		return ecode.TagCreateError
	}
	return
}

func (s *Service) TagUpdate(id int64, params param.TagUpdateParams) (err error) {
	var tag model.Tag
	res := s.Db.First(&tag, id)
	if res.RowsAffected <= 0 {
		err = ecode.TagNotFound
		return
	}

	// 标签不能重名
	if err = s.checkTagName(params.Name, []int64{id}); err != nil {
		return
	}

	updated := s.Db.Model(&tag).Updates(model.Tag{Name: params.Name})
	if updated.RowsAffected <= 0 {
		log.Printf("修改标签失败 %s", updated.Error)
		err = ecode.TagUpdateError
		return
	}

	return
}

func (s *Service) TagList(params param.TagListParams) (*rspModel.TagList, error) {
	var (
		tmodel model.Tag
		tags   []*model.Tag
		total  int64
		field  = []string{"id", "name"}
		offset = (params.Page - 1) * params.Size
	)

	res := s.Db.Model(&tmodel).Offset(int(offset)).Limit(int(params.Size)).Select(field).Find(&tags)
	if res.RowsAffected <= 0 {
		return rspModel.NewTagListResponse([]*model.Tag{}, 0), nil
	}

	s.Db.Model(&tmodel).Count(&total)

	return rspModel.NewTagListResponse(tags, total), nil
}

func (s *Service) TagDetail(params param.TagDetailParams) (detail *rspModel.Tag, err error) {
	if params.ID == 0 {
		err = ecode.TagParamError
		return
	}

	var tag model.Tag
	res := s.Db.First(&tag, params.ID)
	if res.RowsAffected <= 0 {
		err = ecode.TagNotFound
		return
	}

	detail = rspModel.NewTagDetailResponse(&tag)
	return
}

func (s *Service) TagDelete(params param.TagDeleteParams) (err error) {
	if params.ID == 0 {
		err = ecode.TagParamError
		return
	}

	var tag model.Tag
	res := s.Db.Delete(&tag, params.ID)
	if res.RowsAffected <= 0 {
		err = ecode.TagDeleteError
		return
	}

	return
}
