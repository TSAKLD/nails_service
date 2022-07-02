package api

import (
	"encoding/json"
	"nails/entity"
	"nails/service"
	"net/http"
)

type Handler struct {
	s service.UseCase
}

func New(s service.UseCase) *Handler {
	h := Handler{
		s: s,
	}

	return &h
}

func (hdr Handler) AddRecord(w http.ResponseWriter, r *http.Request) {
	var re entity.Record

	err := json.NewDecoder(r.Body).Decode(&re)
	if err != nil {
		json.NewEncoder(w).Encode(error.Error(err))
		return
	}

	err = hdr.s.Insert(re)
	if err != nil {
		json.NewEncoder(w).Encode(error.Error(err))
		return
	}

	json.NewEncoder(w).Encode("Запись добавлена")
}

func (hdr Handler) ShowRecords(w http.ResponseWriter, r *http.Request) {
	var ss []entity.Record

	ss, err := hdr.s.Records()
	if err != nil {
		json.NewEncoder(w).Encode(error.Error(err))
		return
	}

	json.NewEncoder(w).Encode(ss)
}
