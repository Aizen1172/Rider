package model

import "gorm.io/gorm"

type Other struct {
	gorm.Model
	Year int    `json:"year" form:"year" gorm:"comment:年份"`
	Name string `json:"name" form:"name" gorm:"comment:名称"`
	User string `json:"user" form:"user" gorm:"comment:使用者"`
}
