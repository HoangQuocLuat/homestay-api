package types

type UpdateHomestayPath struct {
	Id int64 `uri:"id" binding:"required"`
}

type UpdateHomestayRequest struct {
	ServiceID     int64  `json:"service_id"`
	HostID        int64  `json:"host_id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Location      string `json:"location"`
	Address       string `json:"address"`
	CoverImageURL string `json:"cover_image_url"`
	GalleryImages string `json:"gallery_images"`
	Status        int    `json:"status"`
}

type UpdateHomestayInput struct {
	Path    *UpdateHomestayPath    `json:"path"`
	Request *UpdateHomestayRequest `json:"request"`
}
