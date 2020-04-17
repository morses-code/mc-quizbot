package api

import (
	"github.com/gorilla/mux"
	"log"
	"morses-code/mc-quizbot/store"
	"net/http"
)

type QuizAPI struct {
	dataStore store.DataStore
	Router *mux.Router
}

func CreateAndInitialiseQuizAPI(dataStore store.DataStore) {
	router := mux.NewRouter()
	api := NewQuizAPI(router, dataStore)

	log.Fatal(http.ListenAndServe(":8000", api.Router))
}

func NewQuizAPI(router *mux.Router, dataStore store.DataStore) *QuizAPI {
	api := &QuizAPI{
		dataStore: dataStore,
		Router:    router,
	}

	api.get("/quiz", api.GetQuizHandler)
	api.get("/question/{number}", api.GetQuestionHandler)
	return api
}

func (api *QuizAPI) get(path string, handler http.HandlerFunc) {
	api.Router.HandleFunc(path, handler).Methods(http.MethodGet)
}