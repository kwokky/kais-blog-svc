package model

import (
	"github.com/kwokky/kais-blog-svc/model"
)

type CategoryList struct {
	Categories []*Category `json:"list"`
	Total      int64       `json:"total"`
}

// New
func NewCategoryListResponse(categories []*model.Category, total int64) *CategoryList {
	list := make([]*Category, 0, len(categories))
	for _, category := range categories {
		c:= NewCategoryDetailResponse(category)
		list = append(list, c)
	}

	return &CategoryList{
		Categories: list,
		Total:      total,
	}
}
