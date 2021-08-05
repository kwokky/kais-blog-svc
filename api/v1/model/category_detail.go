package model

import "github.com/kwokky/kais-blog-svc/model"

type Category struct {
	Id        uint   `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

// New
func NewCategoryDetailResponse(category *model.Category) *Category {
	c := &Category{
		Id:   category.ID,
		Name: category.Name,
	}
	if !category.CreatedAt.IsZero() {
		c.CreatedAt = category.CreatedAt.Format("2006-01-02 15:04")
	}
	return c
}
