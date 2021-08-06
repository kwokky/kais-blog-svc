package model

import "github.com/kwokky/kais-blog-svc/model"

type Tag struct {
	Id        uint   `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

// New
func NewTagDetailResponse(tag *model.Tag) *Tag {
	t := &Tag{
		Id:   tag.ID,
		Name: tag.Name,
	}
	if !tag.CreatedAt.IsZero() {
		t.CreatedAt = tag.CreatedAt.Format("2006-01-02 15:04")
	}
	return t
}
