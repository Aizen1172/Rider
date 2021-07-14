package request

import (
	"gorm.io/gorm"
	"riderStyemProject/model"
)

type Other struct {
	Year int    `json:"year"`
	Name string `json:"name"`
	User string `json:"user"`
}

type CreatOther struct {
	Other
}

type UpdateOther struct {
	gorm.Model
	Other
}

type DeleteOther struct {
	gorm.Model
}

func (other *CreatOther) Creat() *model.Other {
	return &model.Other{
		Model: gorm.Model{},
		Year:  other.Year,
		Name:  other.Name,
		User:  other.User,
	}
}

func (other *UpdateOther) Update() *model.Other {
	return &model.Other{
		Model: gorm.Model{},
		Year:  other.Year,
		Name:  other.Name,
		User:  other.User,
	}
}
