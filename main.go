package main

import (
	"log"
	"morses-code/mc-quizbot/api"
	"morses-code/mc-quizbot/mongo"
	"morses-code/mc-quizbot/store"
)

var _ store.Repository = (*QuizAPIStore)(nil)

type QuizAPIStore struct {
	*mongo.Mongo
}

func main() {
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
	api.CreateAndInitialiseQuizAPI(store)

	/*http.HandleFunc("/question", quiz.GetQuestionHandler)
	http.HandleFunc("/answer", quiz.GetAnswerHandler)
	http.HandleFunc("/create", quiz.CreateQuestionHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))*/
}
