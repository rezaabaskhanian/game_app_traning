package entity

type Question struct {
	ID             uint
	Question       string
	PossibleAnswer []string
	CorrectAnswer  string
	Difficultly    string
	CategoryID     uint
}
