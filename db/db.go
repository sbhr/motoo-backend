package motoodb

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sbhr/motoo-backend/model"
)

// MotooDB has methods to get data of motoodb
type MotooDB interface {
	GetAllConversations() []model.Conversation
	GetConversation(id int) model.Conversation
}

type motoo struct {
	db *gorm.DB
}

// New retrun insatnce has MotooDB interface
func New(db *gorm.DB) MotooDB {
	m := &motoo{
		db: db,
	}
	return m
}

// GetAllConversations retrun all data from conversation table
func (m *motoo) GetAllConversations() []model.Conversation {
	cs := []model.Conversation{}
	m.db.Find(&cs)
	return cs
}

// GetConversation retrun data from conversation table
func (m *motoo) GetConversation(id int) model.Conversation {
	c := model.Conversation{}
	m.db.First(&c, id)
	return c
}

// GetResponseByKeyword
// PostConversation
// DeleteConversation
// UpdateConversation
// PostUser
// PostPlaylog
