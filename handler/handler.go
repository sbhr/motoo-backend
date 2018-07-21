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
	GetConversations(w http.ResponseWriter, r *http.Request)
	GetConversationByID(w http.ResponseWriter, r *http.Request)
	PostConversation(w http.ResponseWriter, r *http.Request)
	UpdateConversation(w http.ResponseWriter, r *http.Request)
	DeleteConversation(w http.ResponseWriter, r *http.Request)
	PostUser(w http.ResponseWriter, r *http.Request)
	PostPlaylog(w http.ResponseWriter, r *http.Request)
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

func (h handler) GetConversations(w http.ResponseWriter, r *http.Request) {
	k := r.URL.Query().Get("q")
	var cs []model.Conversation
	var err error
	if k == "" {
		cs, err = h.db.GetAllConversations()
	} else {
		cs, err = h.db.GetConversationByKeyword(k)
	}
	if err != nil {
		code := http.StatusInternalServerError
		http.Error(w, http.StatusText(code), code)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cs)
	return
}

func (h handler) GetConversationByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		code := http.StatusBadRequest
		http.Error(w, http.StatusText(code), code)
		return
	}

	c, err := h.db.GetConversationByID(id)
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

func (h handler) UpdateConversation(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		code := http.StatusBadRequest
		http.Error(w, http.StatusText(code), code)
		return
	}

	var c model.Conversation
	err = json.NewDecoder(r.Body).Decode(&c)
	// Drain and close the body to let the Transport reuse the connection
	io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
	if err != nil {
		code := http.StatusBadRequest
		http.Error(w, err.Error(), code)
		return
	}

	err = h.db.UpdateConversation(id, c)
	if err != nil {
		code := http.StatusInternalServerError
		http.Error(w, http.StatusText(code), code)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	return
}

func (h handler) DeleteConversation(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		code := http.StatusBadRequest
		http.Error(w, http.StatusText(code), code)
		return
	}

	err = h.db.DeleteConversation(id)
	if err != nil {
		code := http.StatusInternalServerError
		http.Error(w, http.StatusText(code), code)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	return
}

func (h handler) PostUser(w http.ResponseWriter, r *http.Request) {
	var u model.User
	err := json.NewDecoder(r.Body).Decode(&u)
	// Drain and close the body to let the Transport reuse the connection
	io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
	if err != nil {
		code := http.StatusBadRequest
		http.Error(w, err.Error(), code)
		return
	}

	err = h.db.PostUser(u)
	if err != nil {
		code := http.StatusInternalServerError
		http.Error(w, http.StatusText(code), code)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	return
}

func (h handler) PostPlaylog(w http.ResponseWriter, r *http.Request) {
	var p model.Playlog
	err := json.NewDecoder(r.Body).Decode(&p)
	// Drain and close the body to let the Transport reuse the connection
	io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
	if err != nil {
		code := http.StatusBadRequest
		http.Error(w, err.Error(), code)
		return
	}

	err = h.db.PostPlaylog(p)
	if err != nil {
		code := http.StatusInternalServerError
		http.Error(w, http.StatusText(code), code)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	return
}
