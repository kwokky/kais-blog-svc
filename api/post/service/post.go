package service

import (
	"github.com/kwokky/kais-blog-svc/library/ecode"
	"github.com/kwokky/kais-blog-svc/model"
	"log"
)

func (s *Service) CreatePost(title, content, author string, categoryId int64) (err error) {
	if err := s.checkCreatePost(title, content, author); err != nil {
		return err
	}

	post := model.Post{Title: title, Content: content, Author: author, CategoryId: categoryId}
	res := s.Db.Create(&post)
	if res.Error != nil {
		log.Printf("新增文章失败 %s", err)
		return ecode.PostCreateError
	}
	return nil
}

func (s *Service) UpdatePost(id int64, title, content, author string, categoryId int64) (err error) {
	var post model.Post
	res := s.Db.First(&post, id)
	if res.RowsAffected <= 0 {
		err = ecode.PostNotFound
		return
	}

	updated := s.Db.Model(&post).Updates(model.Post{Title: title, Content: content, Author: author, CategoryId: categoryId})
	if updated.RowsAffected <= 0 {
		log.Printf("修改文章失败 %s", updated.Error)
		err = ecode.PostUpdateError
		return
	}

	return nil
}

func (s *Service) ListPost(page, size int64, tag, category string) {

}