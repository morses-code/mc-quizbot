package api

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"morses-code/mc-quizbot/store"
	"net/http"
)

type QuizAPI struct {
	dataStore store.DataStore
	Router *mux.Router
}

func CreateAndInitialiseQuizAPI(ctx context.Context, dataStore store.DataStore) {
	router := mux.NewRouter()
	api := NewQuizAPI(ctx, router, dataStore)

	log.Fatal(http.ListenAndServe(":8000", api.Router))
}

func NewQuizAPI(ctx context.Context, router *mux.Router, dataStore store.DataStore) *QuizAPI {
	api := &QuizAPI{
		dataStore: dataStore,
		Router:    router,
	}

	api.get("/quiz", api.GetQuizHandler)
	return api
}

func (api *QuizAPI) get(path string, handler http.HandlerFunc) {
	api.Router.HandleFunc(path, handler).Methods(http.MethodGet)
}