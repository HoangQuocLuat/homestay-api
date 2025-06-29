package types

type Service struct {
	Name          string `json:"name"`  
}

type ServiceGetAllResponse struct {
	ServiceResponse []Service `json:"service_response"`
}
