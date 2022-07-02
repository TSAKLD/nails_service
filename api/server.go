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

func NewServer(s service.UseCase, p string) *Server {
	hand := New(s)
	r := mux.NewRouter()

	r.HandleFunc("/create", hand.AddRecord).Methods(http.MethodPost)
	r.HandleFunc("/records", hand.ShowRecords).Methods(http.MethodGet)

	return &Server{
		port:   p,
		router: r,
	}
}

func (r *Server) ListenAndServe() error {
	fmt.Println("Server is listening...")

	err := http.ListenAndServe(":"+r.port, r.router)
	return err
}
