package main

import (
	"net/http"

	"github.com/sbhr/motoo-backend/router"
)

func main() {
	http.ListenAndServe(":8080", router.New())
}
