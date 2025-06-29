package model

type Room struct {
	Id            int64   `db:"id"`
	HomestayID    int64   `db:"homestay_id"`
	Name          string  `db:"name"`
	Description   string  `db:"description"`
	PricePerNight int     `db:"price_per_night"`
	MaxGuests     int     `db:"max_guests"`
	NumBedrooms   int     `db:"num_bedrooms"`
	NumBathrooms  int     `db:"num_bathrooms"`
	Area          float64 `db:"area"`
	Status        int     `db:"status"`
	CreatedAt     int64   `db:"created_at"`
	UpdatedAt     int64   `db:"updated_at"`
}
