package utils

import (
	"net/http"
)

func WriteError(w http.ResponseWriter, status int, v any) {
	WriteJSON(w, status, v)
}
