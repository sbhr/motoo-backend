package motoodb

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sbhr/motoo-backend/model"
)

// MotooDB has methods to get data of motoodb
type MotooDB interface {
	GetAllConversations() ([]model.Conversation, error)
	GetConversation(id int) (model.Conversation, error)
	GetConversationIncludeKeyword(keyword string) ([]model.Conversation, error)
	PostConversation(keyword, response string) error
	DeleteConversation(id int) error
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
func (m *motoo) GetAllConversations() ([]model.Conversation, error) {
	cs := []model.Conversation{}
	result := m.db.Find(&cs)
	return cs, result.Error
}

// GetConversation retrun data from conversation table
func (m *motoo) GetConversation(id int) (model.Conversation, error) {
	c := model.Conversation{}
	result := m.db.First(&c, id)
	return c, result.Error
}

// GetConversationIncludeKeyword
func (m *motoo) GetConversationIncludeKeyword(keyword string) ([]model.Conversation, error) {
	cs := []model.Conversation{}
	result := m.db.Where("keyword LIKE ?", "%"+keyword+"%").Find(&cs)
	return cs, result.Error
}

// PostConversation
func (m *motoo) PostConversation(keyword, response string) error {
	c := model.Conversation{
		ID:       0,
		Keyword:  keyword,
		Response: response,
	}
	result := m.db.Create(&c)
	return result.Error
}

// DeleteConversation
func (m *motoo) DeleteConversation(id int) error {
	c := model.Conversation{
		ID: id,
	}
	result := m.db.Delete(&c)
	return result.Error
}

// UpdateConversation
// PostUser
// PostPlaylog
