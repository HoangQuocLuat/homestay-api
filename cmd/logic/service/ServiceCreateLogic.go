package service

import (
	"back-end/cmd/database/model"
	"back-end/cmd/svc"
	"back-end/cmd/types"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type ServiceCreateLogic struct {
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	logHelper *log.Helper
}

func NewServiceCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext, logHelper *log.Helper) ServiceCreateLogic {
	return ServiceCreateLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		logHelper: logHelper,
	}
}

func (l *ServiceCreateLogic) ServiceCreate(input *types.CreateServiceRequest) error {
	l.logHelper.Infof("Start creating new service")

	service := &model.Service{
		Name: input.Name,
	}

	err := l.svcCtx.ServiceRepo.CreateService(l.ctx, service)
	if err != nil {
		l.logHelper.Errorf("Failed to create service: %s", err.Error())
		return err
	}

	l.logHelper.Infof("Service created successfully: %+v", service)
	return nil
}
