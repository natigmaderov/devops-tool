package api

import (
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/natigmaderov/devops-tool/service/user"
	"log"
	"net/http"
)

type Server struct {
	addr string
	db   *pgxpool.Pool
}

func NewAPIServe(addr string, db *pgxpool.Pool) *Server {
	return &Server{addr: addr, db: db}
}

func (s *Server) Run() error {
	router := mux.NewRouter()
	subroutine := router.PathPrefix("/api").Subrouter()
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subroutine)
	log.Println("Listening on " + s.addr)
	return http.ListenAndServe(s.addr, router)
}
