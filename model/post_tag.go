package model

import "gorm.io/gorm"

type PostTag struct {
	gorm.Model
	PostId int `gorm:"not null"`
	TagId int `gorm:"not null"`
}
