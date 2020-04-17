package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func (api *QuizAPI) GetQuizHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	result, err := api.dataStore.Backend.GetQuestions(ctx)
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
