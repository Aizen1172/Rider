package model

import "gorm.io/gorm"

type Rider struct {
	gorm.Model
	Year   int    `json:"year" form:"year" gorm:"comment:年份"`
	Name   string `json:"name" form:"name" gorm:"comment:名称"`
	Height int    `json:"height" form:"height" gorm:"comment:身高"`
	Weight int    `json:"weight" form:"weight" gorm:"comment:体重"`
	Punch  int    `json:"punch" form:"punch" gorm:"comment:拳力"`
	Kick   int    `json:"kick" form:"kick" gorm:"comment:踢力"`

	Others []Other `json:"others" gorm:"foreignKey:Year;references:Year"`
}
