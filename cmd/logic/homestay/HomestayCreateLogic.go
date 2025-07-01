package homestay

import (
	"back-end/cmd/database/model"
	"back-end/cmd/svc"
	"back-end/cmd/types"
	"back-end/cmd/utils"
	"context"
)

type HomestayCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomestayCreateLogic {
	return HomestayCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayCreateLogic) HomestayCreate(input *types.CreateHomestayRequest) (*types.Homestay, error) {

	newHomestay := &model.Homestay{
		ServiceID:     input.ServiceID,
		HostID:        input.HostID,
		Name:          input.Name,
		Description:   input.Description,
		Location:      input.Location,
		Address:       input.Address,
		CoverImageURL: input.CoverImageURL,
		GalleryImages: input.GalleryImages,
		Status:        input.Status,
		CreatedAt:     utils.NowInVietnam(),
		UpdatedAt:     utils.NowInVietnam(),
	}

	createdHomestay, err := l.svcCtx.HomestayRepo.CreateHomestay(l.ctx, newHomestay)
	if err != nil {
		return nil, err
	}

	return l.toHomestayModel(createdHomestay), nil
}

func (l *HomestayCreateLogic) toHomestayModel(homestay *model.Homestay) *types.Homestay {
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
