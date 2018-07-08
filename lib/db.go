package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sbhr/motoo-backend/model"
)

// MotooDB has methods to get data of motoodb
type MotooDB interface {
	GetAllConversations() []model.Conversation
}

type motoo struct {
	db *gorm.DB
}

// New retrun insatnce has MotooDB interface
func New(user, password, host, dbName string) (MotooDB, error) {
	db, err := gorm.Open("mysql", user+":"+password+"@tcp("+host+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	m := &motoo{
		db: db,
	}
	if err != nil {
		return nil, err
	}
	return m, nil
}

// GetAllConversations retrun all data from conversation table
func (m *motoo) GetAllConversations() []model.Conversation {
	cs := []model.Conversation{}
	m.db.Find(&cs)
	return cs
}
