package main

import (
	"fmt"
	"os"

	"github.com/sbhr/motoo-backend/lib"
)

func main() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	motooDB, err := db.New(user, password, host, dbName)
	if err != nil {
		fmt.Println("Failed to create db instance: ", err.Error())
		panic(err.Error())
	}
	c := motooDB.GetAllConversations()
	fmt.Println(c)

}
