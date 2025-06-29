package repo

import (
	"back-end/cmd/database/model"
	"context"
)

type ServiceRepo interface {
	GetServices(ctx context.Context, condition *GetCondition) ([]*model.Service, error)
	CreateService(ctx context.Context, service *model.Service) error
	UpdateService(ctx context.Context, service *model.Service) error
}
