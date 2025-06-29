package homestay

import (
	"back-end/cmd/database/model"
	"back-end/cmd/svc"
	"back-end/cmd/types"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type HomestayCreateLogic struct {
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	logHelper *log.Helper
}

func NewHomestayCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext, logHelper *log.Helper) HomestayCreateLogic {
	return HomestayCreateLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		logHelper: logHelper,
	}
}

func (l *HomestayCreateLogic) HomestayCreate(input *types.CreateHomestayRequest) error {
	l.logHelper.Infof("Start process create homestay")

	homestay := &model.Homestay{
		ServiceID:     input.ServiceID,
		HostID:        input.HostID,
		Name:          input.Name,
		Description:   input.Description,
		Location:      input.Location,
		Address:       input.Address,
		CoverImageURL: input.CoverImageURL,
		GalleryImages: input.GalleryImages,
		Status:        input.Status,
	}

	err := l.svcCtx.HomestayRepo.CreateHomestay(l.ctx, homestay)
	if err != nil {
		l.logHelper.Errorf("Failed to create homestay, error: %s", err.Error())
		return err
	}

	l.logHelper.Infof("Homestay created successfully")
	return nil
}
