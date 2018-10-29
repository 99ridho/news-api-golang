package models

type GeneralResponse struct {
	Message      string      `json:"message"`
	Data         interface{} `json:"data,omitempty"`
	ErrorMessage string      `json:"error_message,omitempty"`
}
