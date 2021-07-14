package router

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"riderStyemProject/global"
	"riderStyemProject/model/request"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			fmt.Println("token为空")
			c.Abort()
			return
		}
		//解析token
		claims, err := jwt.ParseWithClaims(token, &request.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			return global.Secret, nil
		})
		if err != nil {
			c.JSON(200,gin.H{"msg":"授权过期！"})
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}
