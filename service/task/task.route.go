package task

import (
	"github.com/gorilla/mux"
	"github.com/sandeshtamanq/jwt/service/auth"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/add-task", auth.ValidateJwt(HandleAddTask)).Methods("POST")
	router.HandleFunc("/get-task", auth.ValidateJwt(HandleGetTask)).Methods("GET")

}
