package homestay

import (
	"back-end/cmd/database/model"
	"back-end/cmd/svc"
	"back-end/cmd/types"
	"context"
	"fmt"
)

type HomestayGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomestayGetLogic {
	return HomestayGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayGetLogic) HomestayGet(input *types.HomestayGetRequest) (*types.Homestay, error) {

	homestay, err := l.svcCtx.HomestayRepo.GetHomestayByID(l.ctx, input.Id)
	if err != nil {
		fmt.Printf("err: ", err)
		return nil, err
	}

	return l.toHomestayModel(homestay), nil
}

func (l *HomestayGetLogic) toHomestayModel(homestay *model.Homestay) *types.Homestay {
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
