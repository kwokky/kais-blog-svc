package service

import (
	"fmt"
	rspModel "github.com/kwokky/kais-blog-svc/api/model"
	"github.com/kwokky/kais-blog-svc/api/param"
	"github.com/kwokky/kais-blog-svc/library/ecode"
	"github.com/kwokky/kais-blog-svc/model"
	"log"
)

func (s *Service) CreatePost(params param.CreatePostParams) (err error) {
	if err := s.checkCreatePost(params); err != nil {
		return err
	}

	post := model.Post{Title: params.Title, Content: params.Content, Author: params.Author, CategoryId: params.CategoryId}
	res := s.Db.Create(&post)
	if res.Error != nil {
		log.Printf("新增文章失败 %s", err)
		return ecode.PostCreateError
	}
	return
}

func (s *Service) UpdatePost(id int64, params param.UpdatePostParams) (err error) {
	var post model.Post
	res := s.Db.First(&post, id)
	if res.RowsAffected <= 0 {
		err = ecode.PostNotFound
		return
	}

	updated := s.Db.Model(&post).Updates(model.Post{Title: params.Title, Content: params.Content, Author: params.Author, CategoryId: params.CategoryId})
	if updated.RowsAffected <= 0 {
		log.Printf("修改文章失败 %s", updated.Error)
		err = ecode.PostUpdateError
		return
	}

	return
}

func (s *Service) ListPost(params param.ListPostParams) (postList *rspModel.PostList, err error) {
	var (
		pmodel     model.Post
		posts      []*model.Post
		total      int64
		ptablename = pmodel.TableName()
		field      = fmt.Sprintf("`%[1]v`.id, `%[1]v`.title, `%[1]v`.content, `%[1]v`.category_id, `%[1]v`.created_at", ptablename)
		offset     = (params.Page - 1) * params.Size
	)

	query := s.Db.Model(&pmodel).Offset(int(offset)).Limit(int(params.Size)).Select(field)

	if params.Tag != "" {
		assocJoinStr := fmt.Sprintf("JOIN `%s` AS assoc ON `assoc`.`post_id` = `%s`.`id`", model.PostTag{}.TableName(), ptablename)
		tagJoinStr := fmt.Sprintf("JOIN `%s` AS tag ON `assoc`.`tag_id` = `tag`.`id`", model.Tag{}.TableName())
		query = query.Joins(assocJoinStr).Joins(tagJoinStr).Where("tag.name = ?", params.Tag)
	}

	if params.Category != "" {
		subQuery := s.Db.Model(&model.Category{}).Select("id").Where("name = ?", params.Category)
		query = query.Where("category_id = (?)", subQuery)
	}

	res := query.Find(&posts)
	if res.RowsAffected == 0 {
		return rspModel.NewPostListResponse([]*model.Post{}, 0), nil
	}

	s.Db.Model(&pmodel).Count(&total)

	return rspModel.NewPostListResponse(posts, total), nil
}

func (s *Service) DetailPost(params param.DetailPostParams) (detail *rspModel.Post, err error) {
	if params.Id == 0 {
		err = ecode.PostParamError
		return
	}

	var post model.Post
	res := s.Db.First(&post, params.Id)
	if res.RowsAffected <= 0 {
		err = ecode.PostNotFound
		return
	}

	return rspModel.NewPostDetailResponse(&post), nil
}

func (s *Service) DeletePost(params param.DeletePostParams) (err error) {
	if params.Id == 0 {
		err = ecode.PostParamError
		return
	}

	var post model.Post
	res := s.Db.Delete(&post, params.Id)
	if res.RowsAffected <= 0 {
		err = ecode.PostDeleteError
		return
	}

	return
}
