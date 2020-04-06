package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"morses-code/mc-quizbot/quiz"
	"net/http"
)

func main() {
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("quiz").Collection("questions")

	collection.InsertOne(ctx, bson.D{
		{Key: "number", Value: "1"},
		{Key: "question", Value: "What is the best day of the week?"},
		{Key: "answer", Value: "Friday"},
	})

	http.HandleFunc("/", quiz.Handler)
	http.HandleFunc("/quiz", quiz.GetQuizHandler)
	http.HandleFunc("/question", quiz.GetQuestionHandler)
	http.HandleFunc("/answer", quiz.GetAnswerHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}