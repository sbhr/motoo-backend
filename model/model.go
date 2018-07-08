package model

type User struct {
	ID     int64
	UserID string
	Name   string
}

type Conversation struct {
	ID       int64
	Keyword  string
	Response string
}

type Playlog struct {
	ID        int64
	UserID    string
	GameName  string
	StartTime int64
	EndTime   int64
	PlayTime  int64
}
