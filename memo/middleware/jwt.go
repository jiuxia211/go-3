package middleware

import (
	"paa/memo/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func JwT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200
		var data interface{}
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 400
		} else {
			claim, err := utils.ParseToken(token)
			if err != nil {
				code = 403 //Token无权限
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = 401 //Token无效
			}
		}
		if code != 200 {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    "Token解析错误",
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
