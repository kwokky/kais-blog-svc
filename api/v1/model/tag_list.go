package model

import (
	"github.com/kwokky/kais-blog-svc/model"
)

type TagList struct {
	Tags  []*Tag `json:"list"`
	Total int64  `json:"total"`
}

// New
func NewTagListResponse(tags []*model.Tag, total int64) *TagList {
	list := make([]*Tag, 0, len(tags))
	for _, category := range tags {
		t := NewTagDetailResponse(category)
		list = append(list, t)
	}

	return &TagList{
		Tags:  list,
		Total: total,
	}
}
