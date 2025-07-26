package entity

type Game struct {
	ID          uint
	CategoryID  uint
	QuestionIDs []uint
	Players     []Player
}

type Player struct {
	ID      uint
	UserID  uint
	GameID  uint
	Scores  uint
	Answers []PlayerAnswer
}

type PlayerAnswer struct {
	ID         uint
	PlayerID   uint
	QuestionID uint
	Choise     PossibleAnswerChoice
}
