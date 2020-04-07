package main

import (
	"log"
	"morses-code/mc-quizbot/quiz"
	"net/http"
)

func main() {
	http.HandleFunc("/quiz", quiz.GetQuizHandler)
	http.HandleFunc("/question", quiz.GetQuestionHandler)
	http.HandleFunc("/answer", quiz.GetAnswerHandler)
	http.HandleFunc("/create", quiz.CreateQuestionHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
