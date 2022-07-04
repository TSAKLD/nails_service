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

func newHandler(s service.UseCase) *Handler {
	h := Handler{
		s: s,
	}

	return &h
}

func (hdr Handler) RecordAdd(w http.ResponseWriter, r *http.Request) {
	var re entity.Record

	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&re)
	if err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}

	re, err = hdr.s.RecordAdd(ctx, re)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError)
		return
	}

	sendResponse(w, re)
}

func (hdr Handler) Records(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ss, err := hdr.s.Records(ctx)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError)
		return
	}

	sendResponse(w, ss)
}
