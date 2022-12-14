package service

import (
	"blog-service/global"
	"blog-service/internal/dao"
	"context"

	otgorm "github.com/eddycjy/opentracing-gorm"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(otgorm.WithContext(ctx, global.DBEngine))
	return svc
}
