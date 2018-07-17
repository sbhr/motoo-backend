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
	GetConversationIncludeKeyword(keyword string) []model.Conversation
	PostConversation(keyword, response string) error
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

// GetConversationIncludeKeyword
func (m *motoo) GetConversationIncludeKeyword(keyword string) []model.Conversation {
	cs := []model.Conversation{}
	m.db.Where("keyword LIKE ?", "%"+keyword+"%").Find(&cs)
	return cs
}

// PostConversation
func (m *motoo) PostConversation(keyword, response string) error {
	c := model.Conversation{
		ID:       0,
		Keyword:  keyword,
		Response: response,
	}
	err := m.db.Create(&c)
	return err.Error
}

// DeleteConversation
// UpdateConversation
// PostUser
// PostPlaylog
