package helper

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func WriteSuccess(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, JSONResponse{
		Success: true,
		Message: message,
	})
}

func WriteSuccessWithData(w http.ResponseWriter, status int, message string, data any) {
	WriteJSON(w, status, JSONResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, JSONResponse{
		Success: false,
		Message: message,
	})
}