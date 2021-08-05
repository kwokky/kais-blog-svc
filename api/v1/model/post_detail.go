package model

import "github.com/kwokky/kais-blog-svc/model"

type Post struct {
	Id         uint   `json:"id,omitempty"`
	Title      string `json:"title,omitempty"`
	Content    string `json:"content,omitempty"`
	Author     string `json:"author,omitempty"`
	CategoryId int64  `json:"category_id,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
}

// New
func NewPostDetailResponse(post *model.Post) *Post {
	p := &Post{
		Id:         post.ID,
		Title:      post.Title,
		Content:    post.Content,
		Author:     post.Author,
		CategoryId: post.CategoryId,
	}
	if !post.CreatedAt.IsZero() {
		p.CreatedAt = post.CreatedAt.Format("2006-01-02 15:04")
	}

	return p
}
