package internal

import (
	"net/http"

	"github.com/nxshock/signaller/internal/models"
)

func errorResponse(w http.ResponseWriter, code models.ApiError, httpCode int, message string) {
	w.Header().Set("Content-Type", "application/json")

	if message != "" {
		code.Message = message
	}

	w.WriteHeader(httpCode)
	w.Write(code.JSON())
}

func NewError(code models.ApiError, message string) *models.ApiError {
	if message != "" {
		code.Message = message
	}

	return &code
}
