package repo

import (
	"back-end/cmd/database/model"
	"context"
)

type RoomRepo interface {
	GetRooms(context.Context, *GetCondition) ([]*model.Room, error)
	CreateRoom(context.Context, *model.Room) error
	UpdateRoom(context.Context, *model.Room) error
}
