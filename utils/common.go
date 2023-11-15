package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
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

// TransformDateToUnix 将2023-09-24T04:45:05+08:00 格式的string类型的时间转化为时间戳int64
func TransformDateToUnix(date string) int64 {
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		log.Fatal(err)
	}
	return t.Unix()
}

// TransformUnixToDate 将时间戳（int64）转化为2023-09-24T04:45:05+08:00（string）
func TransformUnixToDate(date int64) string {
	timeTemplate := "2006-01-02 15:04:05"
	return time.Unix(date, 0).Format(timeTemplate)
}

// Ffmpeg 视频封面截取
func Ffmpeg(videoURL string, frameNum int) ([]byte, error) {
	// 创建一个临时文件来存储输出图像
	outputFile := "output.jpg"

	// 使用 ffmpeg 从视频中获取指定帧并将其输出到临时文件
	cmd := exec.Command("ffmpeg",
		"-i", videoURL,
		"-vf", fmt.Sprintf("select='gte(n,%d)',vflip", frameNum),
		"-vframes", "1",
		outputFile)

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// 读取临时文件的内容到缓冲区
	buf, err := os.ReadFile(outputFile)
	if err != nil {
		log.Fatal(err)
	}

	// 删除临时文件
	err = os.Remove(outputFile)
	if err != nil {
		log.Println("Error removing temporary file:", err)
	}

	return buf, nil
}
