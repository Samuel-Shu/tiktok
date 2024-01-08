package svc

import (
	"follow/internal/config"
	"follow/models"
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
