package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/sbhr/motoo-backend/db"
	"github.com/sbhr/motoo-backend/handler"
	"github.com/sbhr/motoo-backend/router"
	"google.golang.org/appengine"
)

func main() {
	var connectionName, dbName, user, password, protocol string

	if appengine.IsDevAppServer() {
		connectionName = mustGetenv("DB_CONNECTION_NAME")
		user = mustGetenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		protocol = "tcp"
	} else {
		connectionName = mustGetenv("CLOUDSQL_CONNECTION_NAME")
		user = mustGetenv("CLOUDSQL_USER")
		password = os.Getenv("CLOUDSQL_PASSWORD")
		protocol = "cloudsql"
	}
	dbName = mustGetenv("DATABASE_NAME")

	// Connect database
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, protocol, connectionName, dbName))
	if err != nil {
		log.Fatalf("Could not open db: %v", err)
	}
	defer db.Close()

	m := motoodb.New(db)
	h := handler.New(m)

	http.Handle("/", router.New(h))
	// for local
	// http.ListenAndServe(":8080", router.New(h))
	appengine.Main()
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panicf("%s environment variable not set.", k)
	}
	return v
}
