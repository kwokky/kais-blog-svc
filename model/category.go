package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);not null"`
}

func (Category) TableName() string {
	return "categories"
}
