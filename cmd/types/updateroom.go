package types

type UpdateRoomPath struct {
	Id int64 `uri:"room_id" binding:"required"`
}

type UpdateRoomRequest struct {
	HomestayID    int64   `json:"homestay_id"`     // ID của homestay chứa phòng
	Name          string  `json:"name"`            // Tên phòng
	Description   string  `json:"description"`     // Mô tả chi tiết phòng
	PricePerNight int     `json:"price_per_night"` // Giá thuê mỗi đêm
	MaxGuests     int     `json:"max_guests"`      // Số lượng khách tối đa
	NumBedrooms   int     `json:"num_bedrooms"`    // Số phòng ngủ
	NumBathrooms  int     `json:"num_bathrooms"`   // Số phòng tắm
	Area          float64 `json:"area"`            // Diện tích (m²)
	Status        int     `json:"status"`          // Trạng thái phòng (0 = nháp, 1 = hoạt động, ...)
}

type UpdateRoomInput struct {
	Path    *UpdateRoomPath    `json:"path"`
	Request *UpdateRoomRequest `json:"request"`
	
}

type UpdateRoomResponse struct {
}
