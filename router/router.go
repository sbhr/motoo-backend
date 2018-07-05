package router

import (
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	router := mux.NewRouter()

	for _, r := range routes {
		router.Methods(r.methods...).Path(r.path).Handler(r.handler)
	}

	return router
}
