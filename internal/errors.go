package internal

import (
	"net/http"

	"github.com/nxshock/signaller/internal/models"
)

func errorResponse(w http.ResponseWriter, code models.ApiError, httpCode int, messageOverride string) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(httpCode)

	if messageOverride != "" {
		w.Write(models.NewError(code, messageOverride).JSON())
	} else {
		w.Write(code.JSON())
	}
}
