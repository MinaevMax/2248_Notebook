package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSONError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := map[string]string{"error": message}
	json.NewEncoder(w).Encode(resp)
}

func WriteJSONResponse(w http.ResponseWriter, status int, body interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(body)
}