package model

// User has information of discord users
type User struct {
	ID     int    `gorm:"primary_key"`
	UserID string `gorm:"type:varchar(32)"`
	Name   string `gorm:"type:varchar(128)"`
}

// Conversation has data for the bot to talk
type Conversation struct {
	ID       int    `gorm:"primary_key;type:int"`
	Keyword  string `gorm:"unique;type:varchar(512)"`
	Response string `gorm:"type:varchar(2048)"`
}

// Game has information of games
type Game struct {
	ID   int    `gorm:"primary_key"`
	Name string `gorm:"type:varchar(256)"`
}

// Playlog is game play log
type Playlog struct {
	ID        int `gorm:"primary_key"`
	UserID    int `gorm:"type:int"`
	GameID    int `gorm:"type:int"`
	StartTime int `gorm:"type:int"`
	EndTime   int `gorm:"type:int"`
	PlayTime  int `gorm:"type:int"`
}
