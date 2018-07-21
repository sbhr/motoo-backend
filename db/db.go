package motoodb

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sbhr/motoo-backend/model"
)

// MotooDB has methods for CRUD
type MotooDB interface {
	GetAllConversations() ([]model.Conversation, error)
	GetConversationByID(id int) (model.Conversation, error)
	GetConversationByKeyword(keyword string) ([]model.Conversation, error)
	PostConversation(convo model.Conversation) error
	DeleteConversation(id int) error
	UpdateConversation(id int, convo model.Conversation) error
	PostUser(user model.User) error
	PostPlaylog(playlog model.Playlog) error
}

type motoo struct {
	db *gorm.DB
}

// New return instance has MotooDB interface
func New(db *gorm.DB) MotooDB {
	m := &motoo{
		db: db,
	}
	return m
}

// GetAllConversations return all data from conversation table
func (m *motoo) GetAllConversations() ([]model.Conversation, error) {
	cs := []model.Conversation{}
	result := m.db.Find(&cs)
	return cs, result.Error
}

// GetConversation return data from conversation table
func (m *motoo) GetConversationByID(id int) (model.Conversation, error) {
	c := model.Conversation{}
	result := m.db.First(&c, id)
	return c, result.Error
}

// GetConversationByKeyword
func (m *motoo) GetConversationByKeyword(keyword string) ([]model.Conversation, error) {
	cs := []model.Conversation{}
	result := m.db.Where("keyword LIKE ?", "%"+keyword+"%").Find(&cs)
	return cs, result.Error
}

/*
PostConversation
*/
func (m *motoo) PostConversation(convo model.Conversation) error {
	// New Record
	convo.ID = 0
	result := m.db.Create(&convo)
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
func (m *motoo) UpdateConversation(id int, convo model.Conversation) error {
	c := model.Conversation{
		ID: id,
	}
	result := m.db.Model(&c).Updates(convo)
	return result.Error
}

// PostUser
func (m *motoo) PostUser(user model.User) error {
	// New Record
	user.ID = 0
	result := m.db.Create(&user)
	return result.Error
}

// PostPlaylog
func (m *motoo) PostPlaylog(playlog model.Playlog) error {
	// New Record
	playlog.ID = 0
	result := m.db.Create(&playlog)
	return result.Error
}
