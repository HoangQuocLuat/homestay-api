package postgresql

import (
	"back-end/cmd/database/model"
	"back-end/cmd/database/repo"
	"back-end/core/logger"
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type RoomDB struct {
	table               string
	connect             *sqlx.DB
	IgnoreInsertColumns []string
	builder             sq.StatementBuilderType
}

func NewRoomDB() (repo.RoomRepo, error) {
	db, err := sqlx.Open("postgres", "postgres://postgres:524020@localhost:5432/nckh?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &RoomDB{
		table:               "rooms",
		connect:             db,
		IgnoreInsertColumns: []string{"id"},
		builder:             sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}, nil
}

func (r *RoomDB) Close() {
	_ = r.connect.Close()
}

func (r *RoomDB) GetRooms(ctx context.Context, condition *repo.GetCondition) ([]*model.Room, error) {
	ctxLogger := logger.NewContextLog(ctx)

	db := r.builder.Select("*").From(r.table)
	if condition != nil {
		db = db.Where(sq.Like{"name": fmt.Sprintf("%%%s%%", condition.Key)})
	}

	query, args, err := db.ToSql()
	if err != nil {
		ctxLogger.Errorf("Failed to build query: %s", err)
		return nil, err
	}

	var result []*model.Room
	err = r.connect.SelectContext(ctx, &result, query, args...)
	if err != nil {
		ctxLogger.Errorf("Failed to select rooms: %s", err)
		return nil, err
	}
	if len(result) == 0 {
		return nil, sql.ErrNoRows
	}

	return result, nil
}

func (r *RoomDB) CreateRoom(ctx context.Context, room *model.Room) error {
	ctxLogger := logger.NewContextLog(ctx)

	db := r.builder.Insert(r.table).
		Columns(GetListColumn(room, r.IgnoreInsertColumns, []string{})...).
		Values(GetListValues(room, r.IgnoreInsertColumns, []string{})...)

	query, args, err := db.ToSql()
	if err != nil {
		ctxLogger.Errorf("Failed to build insert query: %s", err)
		return err
	}

	_, err = r.connect.ExecContext(ctx, query, args...)
	if err != nil {
		ctxLogger.Errorf("Failed to insert room: %s", err)
		return err
	}
	return nil
}

func (r *RoomDB) UpdateRoom(ctx context.Context, room *model.Room) error {
	ctxLogger := logger.NewContextLog(ctx)

	db := r.builder.Update(r.table).
		Set("homestay_id", room.HomestayID).
		Set("name", room.Name).
		Set("description", room.Description).
		Set("price_per_night", room.PricePerNight).
		Set("max_guests", room.MaxGuests).
		Set("num_bedrooms", room.NumBedrooms).
		Set("num_bathrooms", room.NumBathrooms).
		Set("area", room.Area).
		Set("status", room.Status).
		Set("updated_at", room.UpdatedAt).
		Where(sq.Eq{"id": room.Id})

	query, args, err := db.ToSql()
	if err != nil {
		ctxLogger.Errorf("Failed to build update query: %s", err)
		return err
	}

	_, err = r.connect.ExecContext(ctx, query, args...)
	if err != nil {
		ctxLogger.Errorf("Failed to update room: %s", err)
		return err
	}
	return nil
}
