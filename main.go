package main

import (
	"context"
	"log"
	"morses-code/mc-quizbot/api"
	"morses-code/mc-quizbot/mongo"
	"morses-code/mc-quizbot/store"
	"os"
	"os/signal"
	"syscall"
)

var _ store.Repository = (*QuizAPIStore)(nil)

type QuizAPIStore struct {
	*mongo.Mongo
}

func main() {
	ctx := context.Background()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)


	m := &mongo.Mongo{
		Collection: "questions",
		Database:   "quiz",
		URI:        "mongodb://localhost:27017",
	}

	var err error
	m.Session, err = m.Init()
	if err != nil {
		log.Fatal()
	}

	store := store.DataStore{Backend: QuizAPIStore{m}}
	api.CreateAndInitialiseQuizAPI(ctx, store)

	/*http.HandleFunc("/quiz", quiz.GetQuizHandler)
	http.HandleFunc("/question", quiz.GetQuestionHandler)
	http.HandleFunc("/answer", quiz.GetAnswerHandler)
	http.HandleFunc("/create", quiz.CreateQuestionHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))*/
}
