package main

import (
	"net/http"

	"github.com/nxshock/go-lightning/matrix"
)

func errorResponse(w http.ResponseWriter, code matrix.ApiError, httpCode int, message string) {
	w.Header().Set("Content-Type", "application/json")

	if message != "" {
		code.Message = message
	}

	w.WriteHeader(httpCode)
	w.Write(code.JSON())
}

func NewError(code matrix.ApiError, message string) *matrix.ApiError {
	if message != "" {
		code.Message = message
	}

	return &code
}
