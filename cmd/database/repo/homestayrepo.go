package repo

import (
	"back-end/cmd/database/model"
	"context"
)

type HomestayRepo interface {
	GetHomestays(ctx context.Context, condition *HomestayCondition, page, pageSize int) ([]*model.Homestay, int, error)
	CreateHomestay(ctx context.Context, homestay *model.Homestay) (*model.Homestay, error)
	UpdateHomestay(ctx context.Context, homestay *model.Homestay) (*model.Homestay, error)
	GetHomestayByID(ctx context.Context, id int64) (*model.Homestay, error)
}

type HomestayCondition struct {
	Location  *string
	HostID    *int64
	ServiceID *int64
	Status    *int
}
