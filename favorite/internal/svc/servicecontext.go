package svc

import (
	"favorite/internal/config"
	"favorite/models"
	"gorm.io/gorm"
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
