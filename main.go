package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"morses-code/mc-quizbot/quiz"
	"net/http"
)

func main() {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("quiz").C("questions")
	err = c.Insert(&quiz.Quiz{
		Number:   1,
		Question: "What is the best day of the week?",
		Answer:   "Friday",
	})
	if err != nil {
		log.Fatal(err)
	}

	result := quiz.Quiz{}
	err = c.Find(bson.M{"number" : 1}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Question: ", result.Question)

	http.HandleFunc("/", quiz.Handler)
	http.HandleFunc("/quiz", quiz.GetQuizHandler)
	http.HandleFunc("/question", quiz.GetQuestionHandler)
	http.HandleFunc("/answer", quiz.GetAnswerHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}