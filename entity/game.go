package entity

type Game struct {
	ID          int
	CategoryID  uint
	QuestionIDs []uint

	PlayerIDs []Player
}

type Player struct {
	ID     uint
	UserID uint
	GameID uint
	Score  uint
}
