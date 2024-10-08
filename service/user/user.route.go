package user

import (
	"github.com/gorilla/mux"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/register", HandleRegister).Methods("POST")
	router.HandleFunc("/login", HandleLogin).Methods("POST")

}
