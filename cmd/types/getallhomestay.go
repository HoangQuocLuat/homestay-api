package types

type HomestayGetAllRequest struct {
	PaginationRequest
	Location  *string `form:"location" json:"location"`     // lọc theo location
	HostID    *int64  `form:"host_id" json:"host_id"`       // lọc theo host_id
	ServiceID *int64  `form:"service_id" json:"service_id"` // lọc theo service_id
	Status    *int    `form:"status" json:"status"`         // lọc theo status
}

type HomestayGetAllResponse struct {
	Raw []Homestay `json:"raw"`
	PaginationResponse
}
