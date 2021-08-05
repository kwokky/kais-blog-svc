package service

import (
	"github.com/kwokky/kais-blog-svc/api/v1/param"
	"github.com/kwokky/kais-blog-svc/library/ecode"
	"github.com/kwokky/kais-blog-svc/model"
)

// checkPostCreateParams 检查创建文章参数
func (s *Service) checkPostCreateParams(params param.PostCreateParams) error {
	if params.Title == "" {
		return ecode.PostTitleEmpty
	}
	if params.Content == "" {
		return ecode.PostContentEmpty
	}
	if params.Author == "" {
		return ecode.PostAuthorEmpty
	}

	return nil
}

// checkPostExist 检查文章是否存在
func (s *Service) checkPostExist(id int64) error {
	var post model.Post
	res := s.Db.First(&post, id)
	if res.RowsAffected <= 0 {
		return ecode.PostNotFound
	}
	return nil
}
