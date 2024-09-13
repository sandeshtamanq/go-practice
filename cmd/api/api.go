package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sandeshtamanq/jwt/service/task"
	"github.com/sandeshtamanq/jwt/service/user"
	"gorm.io/gorm"
)

type APIServer struct {
	addr string
	db   *gorm.DB
}

func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		db:   db,
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	taskHandler := task.NewHandler()

	userHandler.RegisterRoutes(subRouter)
	taskHandler.RegisterRoutes(subRouter)

	return http.ListenAndServe(s.addr, router)

}
