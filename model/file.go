package model

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Url  string `json:"url" form:"url" gorm:"comment:文件地址"`
	Name string `json:"name" form:"name" gorm:"comment:文件名"`
	Type string `json:"type" form:"type" gorm:"comment:文件类型"`
}
