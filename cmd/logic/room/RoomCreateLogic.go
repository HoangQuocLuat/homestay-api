package room

import (
	"back-end/cmd/database/model"
	"back-end/cmd/svc"
	"back-end/cmd/types"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type RoomCreateLogic struct {
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	logHelper *log.Helper
}

func NewRoomCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext, logHelper *log.Helper) RoomCreateLogic {
	return RoomCreateLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		logHelper: logHelper,
	}
}

func (l *RoomCreateLogic) RoomCreate(input *types.CreateRoomRequest) error {
	l.logHelper.Infof("Start creating new room")

	room := &model.Room{
		HomestayID:    input.HomestayID,
		Name:          input.Name,
		Description:   input.Description,
		PricePerNight: input.PricePerNight,
		MaxGuests:     input.MaxGuests,
		NumBedrooms:   input.NumBedrooms,
		NumBathrooms:  input.NumBathrooms,
		Area:          input.Area,
		Status:        input.Status,
	}

	err := l.svcCtx.RoomRepo.CreateRoom(l.ctx, room)
	if err != nil {
		l.logHelper.Errorf("Failed to create room: %s", err.Error())
		return err
	}

	l.logHelper.Infof("Room created successfully: %+v", room)
	return nil
}
