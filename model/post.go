package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string `gorm:"type:varchar(100);not null"`
	Content string `gorm:"type:text;not null"`
	CategoryId int64
	Author string `gorm:"type:varchar(100)"`
}
