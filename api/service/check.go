package service

import (
	"github.com/kwokky/kais-blog-svc/api/post/param"
	"github.com/kwokky/kais-blog-svc/library/ecode"
)

func (s *Service) checkCreatePost(params param.CreatePostParams) error {
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
