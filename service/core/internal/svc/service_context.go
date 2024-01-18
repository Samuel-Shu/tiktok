package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
	"mini-tiktok/service/core/internal/config"
	"mini-tiktok/service/core/internal/middleware"
	"mini-tiktok/service/core/models"
	"mini-tiktok/service/favorite/favorite"
	"mini-tiktok/service/follow/follow"
	"mini-tiktok/service/message/message"
)

type ServiceContext struct {
	Config     config.Config
	Engine     *gorm.DB
	RDB        *redis.Client
	Auth       rest.Middleware
	UserModel  *models.DefaultUserModel
	VideoModel *models.DefaultVideoModel
	FollowPb   follow.FollowClient
	FavoritePb favorite.FavoriteClient
	MessagePb  message.MessageClient
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
		FollowPb:   *config.InitFollowClient(c.Etcd.Host),
		FavoritePb: *config.InitFavoriteClient(c.Etcd.Host),
		MessagePb:  *config.InitMessageClient(c.Etcd.Host),
	}
}
