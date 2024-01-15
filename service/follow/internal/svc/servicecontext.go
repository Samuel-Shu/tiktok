package svc

import (
	"mini-tiktok/service/follow/internal/config"
	"mini-tiktok/service/follow/models"
)

type ServiceContext struct {
	Config        config.Config
	RelationModel *models.DefaultRelationModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := models.InitMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		RelationModel: models.NewCommentModel(db),
	}
}
