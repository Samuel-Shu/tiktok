package utils

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"log"
	"tiktok/config"
	"time"
)

type OSSConfig struct {
	AccessKey     string `yaml:"AccessKey"`
	SecretKey     string `yaml:"SecretKey"`
	VideoBucket   string `yaml:"VideoBucket"`
	PictureBucket string `yaml:"PictureBucket"`
	DomainVideo   string `yaml:"DomainVideo"`
	DomainPicture string `yaml:"DomainPicture"`
}

// NewOSSConfig 从*gin.context中获取OSS云配置信息
func NewOSSConfig(c *gin.Context) *OSSConfig {
	serverConfig, exist := c.Get("ServerConfig")
	if !exist {
		log.Fatal("get OssConfig failed !")
	}
	return &OSSConfig{
		AccessKey:     serverConfig.(*config.ServerConfig).AccessKey,
		SecretKey:     serverConfig.(*config.ServerConfig).SecretKey,
		VideoBucket:   serverConfig.(*config.ServerConfig).VideoBucket,
		PictureBucket: serverConfig.(*config.ServerConfig).PictureBucket,
		DomainVideo:   serverConfig.(*config.ServerConfig).DomainVideo,
		DomainPicture: serverConfig.(*config.ServerConfig).DomainPicture,
	}
}

// PushVideo 上传视频至OSS
func (o *OSSConfig) PushVideo(key string, data []byte) bool {

	fmt.Println(o.VideoBucket)

	putPolicy := storage.PutPolicy{
		Scope: o.VideoBucket,
	}
	mac := qbox.NewMac(o.AccessKey, o.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	dataLen := int64(len(data))
	err := formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		fmt.Println("one:", err)
	}
	fmt.Println(ret.Key, ret.Hash)
	if ret.Hash != "" {
		return true
	}
	return false
}

// GetVideo 从OSS获取视频外链
func (o *OSSConfig) GetVideo(key string) string {
	domain := o.DomainVideo
	mac := qbox.NewMac(o.AccessKey, o.SecretKey)
	deadline := time.Now().Add(time.Second * 3600 * 24 * 365).Unix() //1年有效期
	privateAccessURL := storage.MakePrivateURL(mac, domain, key, deadline)
	fmt.Println(privateAccessURL)
	return privateAccessURL
}

// DeleteVideo 将OSS上的指定视频删除
func (o *OSSConfig) DeleteVideo() {
	bucket := o.VideoBucket
	key := "github-x.jpg"
	mac := qbox.NewMac(o.AccessKey, o.SecretKey)
	cfg := storage.Config{}
	bucketManager := storage.NewBucketManager(mac, &cfg)
	err := bucketManager.Delete(bucket, key)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// PushVideoCover 上传视频封面至OSS（上传图片至OSS）
func (o *OSSConfig) PushVideoCover(key string, data []byte) bool {
	putPolicy := storage.PutPolicy{
		Scope: o.PictureBucket,
	}
	mac := qbox.NewMac(o.AccessKey, o.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	dataLen := int64(len(data))
	err := formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret.Key, ret.Hash)
	if ret.Hash != "" {
		return true
	}
	return false
}

// GetCover 获取视频封面链接至OSS（获取图片链接）
func (o *OSSConfig) GetCover(key string) string {
	domain := o.DomainPicture
	mac := qbox.NewMac(o.AccessKey, o.SecretKey)
	deadline := time.Now().Add(time.Second * 3600 * 24 * 365).Unix() //1年有效期
	privateAccessURL := storage.MakePrivateURL(mac, domain, key, deadline)
	return privateAccessURL
}
