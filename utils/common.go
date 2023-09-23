package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
)

// Md5 使用md5加密技术对password进行hash加密
func Md5(password string) string {
	h := md5.New()
	_, err := io.WriteString(h, password)
	if err != nil {
		log.Fatal("md5 failed", err)
	}
	md5Password := string(h.Sum([]byte(nil)))
	md5Password = fmt.Sprintf("%x", md5Password)
	return md5Password
}
