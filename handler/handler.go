package handler

import (
	"net/http"
)

type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}
