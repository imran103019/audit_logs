
package dtos

type ConsumerResponse struct {
	Id            int       `json:"id,omitempty"`
	AppName       string    `json:"app_name"`
	Token         string    `json:"token"`
}

type ConsumerRequest struct {
	AppName        string    `json:"app_name"    form:"app_name"`
}