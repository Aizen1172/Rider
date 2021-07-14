package request

import (
	"gorm.io/gorm"
	"riderStyemProject/model"
)

type User struct {
	Name     string `json:"name" form:"name" gorm:"comment:用户名"`
	Password string `json:"password" form:"password" gorm:"comment:密码"`
}

type CreatUser struct {
	User
}

type UpdateUser struct {
	gorm.Model
	User
}

type DeleteUser struct {
	gorm.Model
}

func (user *CreatUser) Creat() *model.User {
	return &model.User{
		Model:    gorm.Model{},
		Name:     user.Name,
		Password: user.Password,
	}
}

func (user *UpdateUser) Update() *model.User {
	return &model.User{
		Model:    gorm.Model{},
		Name:     user.Name,
		Password: user.Password,
	}
}
