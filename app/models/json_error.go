package models

type JsonError struct {
	Code    int    `json:"error_code"`
	Message string `json:"error_message"`
}
