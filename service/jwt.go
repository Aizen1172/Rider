package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"riderStyemProject/global"
	"riderStyemProject/model"
	"riderStyemProject/model/request"
	"riderStyemProject/model/response"
	"time"
)

var _token string

func Token(user *model.User) (login *response.Login, err error) {
	claims := request.MyClaims{
		ID:   user.ID,
		Name: user.Name,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 604800,
			Issuer:    "Aizen",
		},
	}
	//生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if _token, err = token.SignedString(global.Secret); err != nil {
		return nil, errors.New("创建token失败。")
	}
	login = &response.Login{
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		Token:     _token,
		User:      user,
	}
	return login, err
}
