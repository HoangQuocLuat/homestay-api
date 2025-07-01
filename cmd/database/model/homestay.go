package model

import "time"

type Homestay struct {
	Id            int64     `db:"id"`
	ServiceID     int64     `db:"service_id"`
	HostID        int64     `db:"host_id"`
	Name          string    `db:"name"`
	Description   string    `db:"description"`
	Location      string    `db:"location"`
	Address       string    `db:"address"`
	CoverImageURL string    `db:"cover_image_url"`
	GalleryImages string    `db:"gallery_images"`
	Status        int       `db:"status"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}
