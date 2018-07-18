package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sbhr/motoo-backend/db"
)

// Handler has functions to operate database
type Handler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	db motoodb.MotooDB
}

// New return Handler struct
func New(m motoodb.MotooDB) Handler {
	return handler{
		db: m,
	}
}

func (h handler) GetAll(w http.ResponseWriter, r *http.Request) {
	cs, err := h.db.GetAllConversations()
	if err != nil {
		code := http.StatusInternalServerError
		http.Error(w, err.Error(), code)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cs)
	return
}
