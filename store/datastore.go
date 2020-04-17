package store

import (
	"context"
	"morses-code/mc-quizbot/quiz"
)

type DataStore struct {
	Backend Repository
}

type Repository interface {
	GetQuestions(ctx context.Context) ([]quiz.Question, error)
}
