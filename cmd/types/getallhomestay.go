package types

type Homestay struct {
	ServiceID     int64  `json:"service_id"`      // ID của loại dịch vụ homestay (nếu có)
	HostID        int64  `json:"host_id"`         // Người đăng homestay
	Name          string `json:"name"`            // Tên homestay
	Description   string `json:"description"`     // Mô tả chi tiết
	Location      string `json:"location"`        // Tỉnh/Thành phố
	Address       string `json:"address"`         // Địa chỉ cụ thể
	CoverImageURL string `json:"cover_image_url"` // Ảnh đại diện
	GalleryImages string `json:"gallery_images"`  // Danh sách ảnh (dạng JSON array string)
	Status        int    `json:"status"`          // Trạng thái (ví dụ: 0 = nháp, 1 = hoạt động)
}

type HomestayGetAllResponse struct {
	HomestayResponse []Homestay `json:"homestay_response"`
}
