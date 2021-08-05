package model

import (
	"github.com/kwokky/kais-blog-svc/model"
)

type Post struct {
	Id         uint   `json:"id,omitempty"`
	Title      string `json:"title,omitempty"`
	Content    string `json:"content,omitempty"`
	Author     string `json:"author,omitempty"`
	CategoryId int64  `json:"category_id,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
}

type PostList struct {
	Posts []*Post `json:"list"`
	Total int64   `json:"total"`
}

// New
func NewPostListResponse(posts []*model.Post, total int64) *PostList {
	list := make([]*Post, 0, len(posts))
	for _, post := range posts {
		p := &Post{
			Id:         post.ID,
			Title:      post.Title,
			Content:    post.Content,
			Author:     post.Author,
			CategoryId: post.CategoryId,
			CreatedAt:  post.CreatedAt.Format("2006-01-02 15:04"),
		}
		list = append(list, p)
	}

	return &PostList{
		Posts: list,
		Total: total,
	}
}
