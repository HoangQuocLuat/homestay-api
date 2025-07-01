package postgresql

import (
	"back-end/cmd/database/model"
	"back-end/cmd/database/repo"
	"context"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type HomestayDB struct {
	table               string
	connect             *gorm.DB
	IgnoreInsertColumns []string
}

func NewHomestayDB() (repo.HomestayRepo, error) {
	dsn := "host=localhost user=homestay_api password=123456 dbname=homestay_api port=5435 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &HomestayDB{
		table:               "homestays",
		connect:             db,
		IgnoreInsertColumns: []string{"id"},
	}, nil
}

func (h *HomestayDB) Close() {
	sqlDB, err := h.connect.DB()
	if err == nil {
		_ = sqlDB.Close()
	}
}

func (h *HomestayDB) GetHomestays(ctx context.Context, condition *repo.HomestayCondition, page, pageSize int) ([]*model.Homestay, int, error) {
	db := h.connect.WithContext(ctx).Table(h.table)

	// Apply filter
	if condition != nil {
		if condition.Location != nil {
			db = db.Where("location = ?", condition.Location)
		}
		if condition.HostID != nil {
			db = db.Where("host_id = ?", *condition.HostID)
		}
		if condition.ServiceID != nil {
			db = db.Where("service_id = ?", *condition.ServiceID)
		}
		if condition.Status != nil {
			db = db.Where("status = ?", *condition.Status)
		}
	}

	// Count total
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	if page < 1 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	db = db.Limit(pageSize).Offset(offset)

	var result []*model.Homestay
	if err := db.Find(&result).Error; err != nil {
		return nil, 0, err
	}

	return result, int(total), nil
}

func (h *HomestayDB) CreateHomestay(ctx context.Context, homestay *model.Homestay) (*model.Homestay, error) {
	db := h.connect.WithContext(ctx)
	
	created := *homestay
	if err := db.Table(h.table).Create(&created).Error; err != nil {
		return nil, err
	}

	return &created, nil
}

func (h *HomestayDB) UpdateHomestay(ctx context.Context, homestay *model.Homestay) (*model.Homestay, error) {
	db := h.connect.WithContext(ctx)

	updated := *homestay
	if err := db.Table(h.table).Where("id = ?", homestay.Id).Updates(&updated).Error; err != nil {
		return nil, err
	}

	// Get the updated record
	if err := db.Table(h.table).Where("id = ?", homestay.Id).First(&updated).Error; err != nil {
		return nil, err
	}

	return &updated, nil
}

func (h *HomestayDB) GetHomestayByID(ctx context.Context, id int64) (*model.Homestay, error) {
	db := h.connect.WithContext(ctx)

	var result model.Homestay
	if err := db.Table(h.table).Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}