package store

import (
	"morses-code/mc-quizbot/quiz"
)

type DataStore struct {
	Backend Repository
}

type Repository interface {
	GetQuestions() ([]quiz.Question, error)
	GetQuestion(int) (string, error)
	GetAnswer(int) (string, error)
	CreateQuestion(question quiz.Question) (quiz.Question, error)
}
