package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"tiktok/model"
	"tiktok/utils"
)

// UserRegister :用户注册
func UserRegister(c *gin.Context) {
	type RegisterStatus struct {
		model.HttpStatus
		UserId int64
		Token  string
	}
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username)
	status := model.UserRegister(username, utils.Md5(password))

	if status != 0 {
		c.JSON(http.StatusOK, RegisterStatus{
			HttpStatus: model.HttpStatus{
				StatusCode: 1,
				StatusMsg:  "register failed",
			},
		})
		return
	}
	c.JSON(http.StatusOK, RegisterStatus{
		HttpStatus: model.HttpStatus{
			StatusCode: 0,
			StatusMsg:  "register successful",
		},
		UserId: model.FindUser(username),
		Token:  username + password,
	})

}

// UserLogin 用户登录
func UserLogin(c *gin.Context) {
	type RegisterStatus struct {
		model.HttpStatus
		UserId int64
		Token  string
	}
	name := c.PostForm("username")
	password := c.PostForm("password")

	if !model.Login(name, utils.Md5(password)) {
		c.JSON(http.StatusOK, RegisterStatus{
			HttpStatus: model.HttpStatus{
				StatusCode: 1,
				StatusMsg:  "login failed!",
			}})
		return
	}
	c.JSON(http.StatusOK, RegisterStatus{
		HttpStatus: model.HttpStatus{
			StatusCode: 0,
			StatusMsg:  "login successful",
		},
		UserId: model.FindUser(name),
		Token:  name + password,
	})

}

// GetUserInfo todo:获取用户信息
func GetUserInfo(c *gin.Context) {
	type UserInfo struct {
		model.HttpStatus
		User model.User `json:"user"`
	}

	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		log.Fatal("translate userId(string) to userId(int) is failed !", err)
	}
	//用户不存在，返回错误信息
	if !model.FindUserWithId(int64(userId)) {
		c.JSON(http.StatusOK, UserInfo{HttpStatus: model.HttpStatus{
			StatusCode: 1,
			StatusMsg:  "the user is not exist !",
		}})
		return
	}

	c.JSON(http.StatusOK, UserInfo{
		HttpStatus: model.HttpStatus{
			StatusCode: 0,
			StatusMsg:  "get userInfo successful",
		},
		User: model.GetUserData(int64(userId)),
	})
}
