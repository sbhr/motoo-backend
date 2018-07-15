package model

// User has information of discord users
type User struct {
	ID     int64  `gorm:"primary_key"`
	UserID string `gorm:"type:varchar(32)"`
	Name   string `gorm:"type:varchar(128)"`
}

// Conversation has data for the bot to talk
type Conversation struct {
	ID       int64  `gorm:"primary_key;type:int"`
	Keyword  string `gorm:"unique;type:varchar(512)"`
	Response string `gorm:"type:varchar(2048)"`
}

// Playlog is game play log
type Playlog struct {
	ID        int64  `gorm:"primary_key"`
	UserID    string `gorm:"type:int"`
	GameName  string `gorm:"unique;type:varchar(128)"`
	StartTime int64  `gorm:"type:int"`
	EndTime   int64  `gorm:"type:int"`
	PlayTime  int64  `gorm:"type:int"`
}
