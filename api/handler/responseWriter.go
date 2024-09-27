package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type jsonError struct {
	Message string `json:"error"`
}

func newJSONError(message string) *jsonError {
	return &jsonError{
		Message: message,
	}
}

func RespondJSON(w http.ResponseWriter, content any, status int) {
	res, err := json.Marshal(content)
	if err != nil {
		log.Printf("Failed to marshal JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	w.Write(res)
}

func RespondOK(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	w.Write([]byte(`{"ok":true}`))
}

func RespondError(w http.ResponseWriter, message string, status int) {
	RespondJSON(w, newJSONError(message), status)
}
