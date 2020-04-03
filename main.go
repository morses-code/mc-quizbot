package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/question", getQuestionHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func getQuestionHandler(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(quiz)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Fprintf(w, "%s\n", data)
}

type Quiz struct {
	Number int `json:"number"`
	Question string `json:"question"`
	Answer string `json:"answer"`
}

var quiz = []Quiz{
	{
		Number:1,
		Question:"What is the best day of the week?",
		Answer:"Friday",
	},
}
