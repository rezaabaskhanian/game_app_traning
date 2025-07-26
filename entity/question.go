package entity

type Question struct {
	ID              uint
	Text            string
	PossibleAnswers []PossibleAnswers
	CorrectAnswerID uint
	Difficultly     string
	CategoryID      uint
}

type PossibleAnswers struct {
	ID     uint
	Text   string
	Choise uint8
}

type PossibleAnswerChoice uint

const (
	PossibleAnswerA PossibleAnswerChoice = iota + 1
	PossibleAnswerB
	PossibleAnswerC
	PossibleAnswerD
)
