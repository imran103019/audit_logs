package dtos
type SuccessResponse struct {
	Success bool `json:"success"`
}
type UnauthorizedResponse struct {
	Message string `json:"message"`
}

type Response struct {
	Data        interface{}  `json:"data,omitempty"`
	Pagination  interface{}  `json:"pagination,omitempty"`
}
