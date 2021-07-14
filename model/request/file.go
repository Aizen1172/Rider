package request

import (
	"gorm.io/gorm"
	"riderStyemProject/model"
)

type File struct {
	Url  string `json:"url" form:"url" gorm:"comment:文件地址"`
	Name string `json:"name" form:"name" gorm:"comment:文件名"`
	Type string `json:"type" form:"type" gorm:"comment:文件类型"`
}

type DeleteFile struct {
	gorm.Model
	File
}

func (file *File) Creat() *model.File {
	return &model.File{
		Url:  file.Url,
		Name: file.Name,
		Type: file.Type,
	}
}

func (file *File)Update() *model.File {
	return &model.File{
		Url:   file.Url,
		Name:  file.Name,
		Type:  file.Type,
	}
}