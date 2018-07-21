package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/sbhr/motoo-backend/db"
	"github.com/sbhr/motoo-backend/handler"
	"github.com/sbhr/motoo-backend/router"
)

func main() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// Connect database
	db, err := gorm.Open("mysql", user+":"+password+"@tcp("+host+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Failed to create db instance: ", err.Error())
		panic(err.Error())
	}
	defer db.Close()

	m := motoodb.New(db)
	if err != nil {
		fmt.Println("Failed to create db instance: ", err.Error())
		panic(err.Error())
	}

	h := handler.New(m)

	http.ListenAndServe(":8080", router.New(h))
}
