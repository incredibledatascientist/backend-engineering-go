package utils

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
	Error   any    `json:"error,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, resp APIResponse) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)
}

func Success(w http.ResponseWriter, status int, message string, data any) {
	resp := APIResponse{
		Success: true,
		Data:    data,
		Message: message,
		Error:   nil,
	}
	writeJSON(w, status, resp)
}

func Error(w http.ResponseWriter, status int, message string, err any) {
	resp := APIResponse{
		Success: false,
		Data:    nil,
		Message: message,
		Error:   err,
	}
	writeJSON(w, status, resp)
}
