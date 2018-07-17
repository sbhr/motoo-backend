package motoodb

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sbhr/motoo-backend/model"
)

// MotooDB has methods for CRUD
type MotooDB interface {
	GetAllConversations() ([]model.Conversation, error)
	GetConversation(id int) (model.Conversation, error)
	GetConversationIncludeKeyword(keyword string) ([]model.Conversation, error)
	PostConversation(keyword, response string) error
	DeleteConversation(id int) error
	UpdateConversation(id int, keyword, response string) error
	PostUser(userID, name string) error
	PostPlaylog(userID, gameName string, startTime, endTime, playTime int) error
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

/*
PostConversation

	aaaa
*/
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
func (m *motoo) UpdateConversation(id int, keyword, response string) error {
	c := model.Conversation{
		ID: id,
	}
	result := m.db.Model(&c).Updates(model.Conversation{Keyword: keyword, Response: response})
	return result.Error
}

// PostUser
func (m *motoo) PostUser(userID, name string) error {
	u := model.User{
		ID:     0,
		UserID: userID,
		Name:   name,
	}
	result := m.db.Create(&u)
	return result.Error
}

// PostPlaylog
func (m *motoo) PostPlaylog(userID, gameName string, startTime, endTime, playTime int) error {
	p := model.Playlog{
		ID:        0,
		UserID:    userID,
		GameName:  gameName,
		StartTime: startTime,
		EndTime:   endTime,
		PlayTime:  playTime,
	}
	result := m.db.Create(&p)
	return result.Error
}
