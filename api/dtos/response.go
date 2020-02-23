package dtos
type SuccessResponse struct {
	Success bool `json:"success"`
}

type Response struct {
	Data        interface{}  `json:"data,omitempty"`
	Pagination  interface{}  `json:"pagination,omitempty"`
}
