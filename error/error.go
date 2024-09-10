package error

import (
	"net/http"

	"github.com/sandeshtamanq/jwt/utils"
)

func WriteError(w http.ResponseWriter, status int, v any) {
	utils.WriteJSON(w, status, v)
}
