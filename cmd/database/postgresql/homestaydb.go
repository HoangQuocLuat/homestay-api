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

type HomestayDB struct {
	table               string
	connect             *sqlx.DB
	IgnoreInsertColumns []string
	builder             sq.StatementBuilderType
}

func NewHomestayDB() (repo.HomestayRepo, error) {
	db, err := sqlx.Open("postgres", "postgres://postgres:524020@localhost:5432/nckh?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &HomestayDB{
		table:               "homestays",
		connect:             db,
		IgnoreInsertColumns: []string{"id"},
		builder:             sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}, nil
}

func (h *HomestayDB) Close() {
	_ = h.connect.Close()
}

func (h *HomestayDB) GetHomestays(ctx context.Context, condition *repo.GetCondition) ([]*model.Homestay, error) {
	ctxLogger := logger.NewContextLog(ctx)

	db := h.builder.Select("*").From(h.table)
	if condition != nil {
		db = db.Where(sq.Like{"name": fmt.Sprintf("%%%s%%", condition.Key)})
	}

	query, args, err := db.ToSql()
	if err != nil {
		ctxLogger.Errorf("Failed to build query: %s", err)
		return nil, err
	}

	var result []*model.Homestay
	err = h.connect.SelectContext(ctx, &result, query, args...)
	if err != nil {
		ctxLogger.Errorf("Failed to select homestays: %s", err)
		return nil, err
	}
	if len(result) == 0 {
		return nil, sql.ErrNoRows
	}

	return result, nil
}

func (h *HomestayDB) CreateHomestay(ctx context.Context, homestay *model.Homestay) error {
	ctxLogger := logger.NewContextLog(ctx)

	db := h.builder.Insert(h.table).
		Columns(GetListColumn(homestay, h.IgnoreInsertColumns, []string{})...).
		Values(GetListValues(homestay, h.IgnoreInsertColumns, []string{})...)

	query, args, err := db.ToSql()
	if err != nil {
		ctxLogger.Errorf("Failed to build insert query: %s", err)
		return err
	}

	_, err = h.connect.ExecContext(ctx, query, args...)
	if err != nil {
		ctxLogger.Errorf("Failed to insert homestay: %s", err)
		return err
	}
	return nil
}

func (h *HomestayDB) UpdateHomestay(ctx context.Context, homestay *model.Homestay) error {
	ctxLogger := logger.NewContextLog(ctx)

	db := h.builder.Update(h.table).
		Set("service_id", homestay.ServiceID).
		Set("host_id", homestay.HostID).
		Set("name", homestay.Name).
		Set("description", homestay.Description).
		Set("location", homestay.Location).
		Set("address", homestay.Address).
		Set("cover_image_url", homestay.CoverImageURL).
		Set("gallery_images", homestay.GalleryImages).
		Set("status", homestay.Status).
		Set("updated_at", homestay.UpdatedAt).
		Where(sq.Eq{"id": homestay.Id})

	query, args, err := db.ToSql()
	if err != nil {
		ctxLogger.Errorf("Failed to build update query: %s", err)
		return err
	}

	_, err = h.connect.ExecContext(ctx, query, args...)
	if err != nil {
		ctxLogger.Errorf("Failed to update homestay: %s", err)
		return err
	}
	return nil
}
