package handler

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/sbhr/motoo-backend/db"
	"github.com/sbhr/motoo-backend/model"
)

// Handler has functions to operate database
type Handler interface {
	GetAllConversations(w http.ResponseWriter, r *http.Request)
	GetConversation(w http.ResponseWriter, r *http.Request)
	PostConversation(w http.ResponseWriter, r *http.Request)
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

func (h handler) GetAllConversations(w http.ResponseWriter, r *http.Request) {
	cs, err := h.db.GetAllConversations()
	if err != nil {
		code := http.StatusInternalServerError
		http.Error(w, http.StatusText(code), code)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cs)
	return
}

func (h handler) GetConversation(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		code := http.StatusBadRequest
		http.Error(w, http.StatusText(code), code)
		return
	}

	c, err := h.db.GetConversation(id)
	if err != nil {
		code := http.StatusInternalServerError
		http.Error(w, http.StatusText(code), code)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
	return
}

func (h handler) PostConversation(w http.ResponseWriter, r *http.Request) {
	var c model.Conversation
	err := json.NewDecoder(r.Body).Decode(&c)
	// Drain and close the body to let the Transport reuse the connection
	io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
	if err != nil {
		code := http.StatusBadRequest
		http.Error(w, err.Error(), code)
		return
	}

	err = h.db.PostConversation(c)
	if err != nil {
		code := http.StatusInternalServerError
		http.Error(w, http.StatusText(code), code)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	return
}
