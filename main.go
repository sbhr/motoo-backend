package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/sbhr/motoo-backend/db"
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

	motooDB := motoodb.New(db)
	if err != nil {
		fmt.Println("Failed to create db instance: ", err.Error())
		panic(err.Error())
	}

	cs := motooDB.GetAllConversations()
	fmt.Println(cs)

	fmt.Println("============")

	c := motooDB.GetConversation(3)
	fmt.Println(c)
}
