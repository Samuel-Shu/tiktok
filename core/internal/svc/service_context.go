package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
	"mini-tiktok/core/internal/config"
	"mini-tiktok/core/internal/middleware"
	"mini-tiktok/core/models"
)

type ServiceContext struct {
	Config config.Config
	Engine *gorm.DB
	RDB    *redis.Client
	Auth   rest.Middleware

	UserModel  *models.DefaultUserModel
	VideoModel *models.DefaultVideoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	engine := models.InitMysql(c.Mysql.DataSource)
	engine.Logger.LogMode(4)
	return &ServiceContext{
		Config:     c,
		Engine:     engine,
		RDB:        models.InitRedis(c),
		UserModel:  models.NewUserModel(engine),
		VideoModel: models.NewVideoModel(engine),
		Auth:       middleware.NewAuthMiddleware().Handle,
	}
}
