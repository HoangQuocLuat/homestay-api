package homestay

import (
	"back-end/cmd/database/model"
	"back-end/cmd/svc"
	"back-end/cmd/types"
	"context"
	"time"
)

type HomestayUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomestayUpdateLogic {
	return HomestayUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayUpdateLogic) HomestayUpdate(input *types.UpdateHomestayInput) (*types.Homestay, error) {

	homestay := &model.Homestay{
		Id:            input.Path.Id,
		ServiceID:     input.Request.ServiceID,
		HostID:        input.Request.HostID,
		Name:          input.Request.Name,
		Description:   input.Request.Description,
		Location:      input.Request.Location,
		Address:       input.Request.Address,
		CoverImageURL: input.Request.CoverImageURL,
		GalleryImages: input.Request.GalleryImages,
		Status:        input.Request.Status,
		UpdatedAt:     time.Now(),
	}

	updatedHomestay, err := l.svcCtx.HomestayRepo.UpdateHomestay(l.ctx, homestay)
	if err != nil {
		return nil, err
	}

	return l.toHomestayModel(updatedHomestay), nil
}

func (l *HomestayUpdateLogic) toHomestayModel(homestay *model.Homestay) *types.Homestay {
	return &types.Homestay{
		Id:            homestay.Id,
		ServiceID:     homestay.ServiceID,
		HostID:        homestay.HostID,
		Name:          homestay.Name,
		Description:   homestay.Description,
		Location:      homestay.Location,
		Address:       homestay.Address,
		CoverImageURL: homestay.CoverImageURL,
		GalleryImages: homestay.GalleryImages,
		Status:        homestay.Status,
		CreatedAt:     homestay.CreatedAt,
		UpdatedAt:     &homestay.UpdatedAt,
	}
}
