package request

import (
	"gorm.io/gorm"
	"riderStyemProject/model"
)

type Rider struct {
	Year   int    `json:"year"`
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Punch  int    `json:"punch" `
	Kick   int    `json:"kick"`
}

type CreatRider struct {
	Rider
}

type UpdateRider struct {
	gorm.Model
	Rider
}

type DeleteRider struct {
	gorm.Model
}

func (rider *CreatRider) Creat() *model.Rider {
	return &model.Rider{
		Model:  gorm.Model{},
		Year:   rider.Year,
		Name:   rider.Name,
		Height: rider.Height,
		Weight: rider.Weight,
		Punch:  rider.Punch,
		Kick:   rider.Kick,
	}
}

func (rider *UpdateRider) Update() *model.Rider {
	return &model.Rider{
		Model:  gorm.Model{},
		Year:   rider.Year,
		Name:   rider.Name,
		Height: rider.Height,
		Weight: rider.Weight,
		Punch:  rider.Punch,
		Kick:   rider.Kick,
	}
}
