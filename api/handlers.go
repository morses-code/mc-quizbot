package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func (api *QuizAPI) GetQuizHandler(w http.ResponseWriter, r *http.Request) {
	result, err := api.dataStore.Backend.GetQuestions()
	if err != nil {
		log.Panic(err)
		return
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (api *QuizAPI) GetQuestionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number, err := strconv.Atoi(vars["number"])
	if err != nil {
		log.Panic(err)
	}

	result, err := api.dataStore.Backend.GetQuestion(number)
	if err != nil {
		log.Panic(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (api *QuizAPI) GetAnswerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number, err := strconv.Atoi(vars["number"])
	if err != nil {
		log.Panic(err)
	}

	result, err := api.dataStore.Backend.GetAnswer(number)
	if err != nil {
		log.Panic(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
