package svc

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"mini-tiktok/core/internal/config"
	"mini-tiktok/core/models"
)

type ServiceContext struct {
	Config config.Config
	Engine *gorm.DB
	RDB    *redis.Client

	UserModel *models.DefaultUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	engine := models.InitMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:    c,
		Engine:    engine,
		RDB:       models.InitRedis(c),
		UserModel: models.NewUserModel(engine),
	}

}
