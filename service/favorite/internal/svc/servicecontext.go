package svc

import (
	"gorm.io/gorm"
	"mini-tiktok/service/favorite/internal/config"
	"mini-tiktok/service/favorite/models"
)

type ServiceContext struct {
	Config        config.Config
	Mysql         *gorm.DB
	FavoriteModel *models.DefaultFavoriteModel
	CommentModel  *models.DefaultCommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := models.InitMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:        c,
		Mysql:         db,
		FavoriteModel: models.NewUserModel(db),
		CommentModel:  models.NewCommentModel(db),
	}
}
