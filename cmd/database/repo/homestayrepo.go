package repo

import (
	"back-end/cmd/database/model"
	"context"
)

type HomestayRepo interface {
	GetHomestays(context.Context, *GetCondition) ([]*model.Homestay, error)
	CreateHomestay(context.Context, *model.Homestay) error
	UpdateHomestay(context.Context, *model.Homestay) error
}
