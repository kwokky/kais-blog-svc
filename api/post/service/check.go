package service

import "github.com/kwokky/kais-blog-svc/library/ecode"

func (s *Service) checkCreatePost(title, content, author string) error {
	if title == "" {
		return ecode.PostTitleEmpty
	}
	if content == "" {
		return ecode.PostContentEmpty
	}
	if author == "" {
		return ecode.PostAuthorEmpty
	}

	return nil
}
