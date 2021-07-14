package request

type Login struct {
	Name     string `json:"name" form:"name" gorm:"comment:用户名"`
	Password string `json:"password" form:"password" gorm:"comment:密码"`
	Yzm      string `json:"yzm" form:"yzm" gorm:"comment:验证码"`
	YzmId    string `gorm:"comment:验证码ID"`
}