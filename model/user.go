package model

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"tiktok/db"
	"time"
)

type User struct {
	Avatar          string `json:"avatar"`           // 用户头像
	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
	FavoriteCount   int64  `json:"favorite_count"`   // 喜欢数
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	ID              int64  `json:"id"`               // 用户id
	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
	Name            string `json:"name"`             // 用户名称
	Signature       string `json:"signature"`        // 个人简介
	TotalFavorited  string `json:"total_favorited"`  // 获赞数量
	WorkCount       int64  `json:"work_count"`       // 作品数
}

type tbUser struct {
	User
	Password string `json:"password"`
}

// UserRegister 用户注册
func UserRegister(username, password string) int64 {

	if FindUser(username) == 0 {
		db.Db.Mysql.Model(&tbUser{}).Create(map[string]interface{}{"name": username, "password": password})
		fmt.Println("用户创建成功")
		return 0
	}
	fmt.Println("用户创建失败")
	return 1
}

// FindUser 通过username查找用户并返回其id
func FindUser(username string) int64 {
	tbUser := tbUser{}
	db.Db.Mysql.Where("name = ?", username).First(&tbUser)
	return tbUser.ID
}

// FindUserWithId 通过用户id判断用户是否存在
func FindUserWithId(id int64) bool {
	user := tbUser{}
	db.Db.Mysql.Where("id = ?", id).Find(&user).Limit(1)
	if user.Name == "" {
		return false
	}
	return true
}

// GetUserData 通过传入的id从数据库获取用户信息
func GetUserData(id int64) User {
	user := tbUser{}
	db.Db.Mysql.Where("id=?", id).First(&user)
	return user.User
}

// Login 用户登录
func Login(username, password string) bool {
	userData := tbUser{}
	var res bool
	if FindUser(username) == 0 {
		res = false
	} else {
		db.Db.Mysql.Where("name=?", username).First(&userData)
		if userData.Password == password {
			res = true
		} else {
			res = false
		}
	}
	return res
}

// TransformDateToUnixTest 一个测试函数，将数据库读取的date类型时间转化为时间戳
func TransformDateToUnixTest() {
	type tbUser struct {
		CreateTime string
	}

	date := tbUser{}
	db.Db.Mysql.Where("id = ?", 1).First(&date)
	fmt.Println(date.CreateTime)
	Time, err := time.Parse(time.RFC3339, date.CreateTime)
	if err != nil {
		log.Fatal(err)
	}
	println(Time.Unix())
}

// AddWorkCount 增加作品数量，默认为：1
func AddWorkCount(userid int64) {
	db.Db.Mysql.Model(&tbUser{}).Where("id = ?", userid).Update("work_count", gorm.Expr("work_count + ?", 1))
}

func UpdateWorkCount(userid, count int64) {
	db.Db.Mysql.Model(&tbUser{}).Where("id = ?", userid).Update("work_count", count)
}
