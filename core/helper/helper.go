package helper

import (
	"crypto/md5"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"mini-tiktok/core/define"
	"time"
)

func GenerateToken(id int64, username string, second int) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(second)).Unix(),
		},
	}
	//fmt.Printf("uc: %+v", uc)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte("mini-tiktok"))
	if err != nil {
		log.Println("err:", err)
		return "", err
	}

	//tokenString := fmt.Sprintf("%v", token)
	return tokenString, nil
}

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
