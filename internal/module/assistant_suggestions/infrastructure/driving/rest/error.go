package rest

import (
	"common"
	"encoding/json"
	"net/http"
)

type apiError struct {
	Error struct{
		Code string
		Message string
		RequestID string
	}
}

func writeError(w http.ResponseWriter, r *http.Request, status int, code, message string) {
	
	var body apiError
	body.Error.Code = code
	body.Error.Message = message
	if reqID := r.Header.Get("X-Request-ID"); reqID != "" {
		body.Error.RequestID= reqID
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func writeDomainError(w http.ResponseWriter, r *http.Request, err error) {
	if common.IsValidationError(err) {
		writeError(w, r, http.StatusBadRequest, "Bad_request", err.Error())
		return
	}
	if common.IsNotFoundError(err) {
		writeError(w, r, http.StatusNotFound, "not_found", err.Error())
		return
	}
	writeError(w, r, http.StatusInternalServerError, "Internal_error", "internal_server_error")
}