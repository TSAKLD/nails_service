package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"nails/service"
	"net/http"
)

type Server struct {
	port   string
	router *mux.Router
}

func NewServer(p string) *Server {
	return &Server{
		port:   p,
		router: mux.NewRouter(),
	}
}

func (s *Server) setRoutes(us service.UseCase) {
	hand := newHandler(us)

	s.router.HandleFunc("/create", hand.RecordAdd).Methods(http.MethodPost)
	s.router.HandleFunc("/records", hand.Records).Methods(http.MethodGet)
}

func (s *Server) Start(us service.UseCase) error {
	s.setRoutes(us)

	fmt.Println("Server is listening...")

	return http.ListenAndServe(":"+s.port, s.router)
}
