package middleware

import (
	"TikTok/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Information struct {
	UserId   uint64 `json:"user_id"`
	UserName string `json:"username"`
}

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Query("token")
		if tokenStr == "" {
			tokenStr = c.PostForm("token")
		}
		//用户不存在
		if tokenStr == "" {
			c.JSON(http.StatusOK, utils.Response{StatusCode: 401, StatusMsg: "用户不存在"})
			c.Abort() //阻止执行
			return
		}
		//验证token
		tokenStruck, ok := CheckToken(tokenStr)
		if !ok {
			c.JSON(http.StatusOK, utils.Response{
				StatusCode: 403,
				StatusMsg:  "token不正确",
			})
			c.Abort() //阻止执行
			return
		}

		c.Set("username", tokenStruck.UserName)
		c.Set("user_id", tokenStruck.UserId)

		c.Next()
	}
}

func CheckToken(str string) (information Information, bool bool) {
	if str == "123456" {
		information.UserId = 1
		information.UserName = "张三丰"
		bool = true
	} else {
		bool = false
	}
	return information, bool
}
