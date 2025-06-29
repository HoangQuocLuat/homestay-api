package model

type Service struct {
	Id            int64   `db:"id"`
	Name          string  `db:"name"`
	CreatedAt     int64   `db:"created_at"`
	UpdatedAt     int64   `db:"updated_at"`
}
