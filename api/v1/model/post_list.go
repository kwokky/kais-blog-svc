package model

import (
	"github.com/kwokky/kais-blog-svc/model"
)

type PostList struct {
	Posts []*Post `json:"list"`
	Total int64   `json:"total"`
}

// New
func NewPostListResponse(posts []*model.Post, total int64) *PostList {
	list := make([]*Post, 0, len(posts))
	for _, post := range posts {
		p := NewPostDetailResponse(post)
		list = append(list, p)
	}

	return &PostList{
		Posts: list,
		Total: total,
	}
}
