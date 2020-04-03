package quiz

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Quiz struct {
	Number   int    `json:"number"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

var quizData = []Quiz{
	{
		Number:   1,
		Question: "What is the best day of the week?",
		Answer:   "Friday",
	},
	{
		Number:   2,
		Question: "What is the worst day of the week?",
		Answer:   "Monday",
	},
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func GetQuizHandler(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(quizData)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Fprintf(w, "%s\n", data)
}

func GetQuestionHandler(w http.ResponseWriter, r *http.Request) {
	number, err := checkNumber(r)

	question, err := getQuestion(number)
	if err != nil {
		log.Panic(err)
	}

	fmt.Fprintf(w, "%s\n", question)
}

func GetAnswerHandler(w http.ResponseWriter, r *http.Request) {
	number, err := checkNumber(r)

	answer, err := getAnswer(number)
	if err != nil {
		log.Panic(err)
	}

	fmt.Fprintf(w, "%s\n", answer)
}

func checkNumber(r *http.Request) (int, error) {
	numbers, ok := r.URL.Query()["number"]
	if !ok || len(numbers[0]) < 1 {
		log.Panic("no question number provided")
	}

	number, err := strconv.Atoi(numbers[0])
	if err != nil {
		log.Panic("unable to convert string, check you entered a valid numeric value.")
	}

	return number, err
}

func getQuestion(number int) (string, error) {
	result, s, err := getQuiz()
	if err != nil {
		return s, err
	}

	var question string
	for i := range result {
		if result[i].Number == number {
			question = result[i].Question
		}
	}

	return question, nil
}

func getAnswer(number int) (string, error) {
	result, s, err := getQuiz()
	if err != nil {
		return s, err
	}

	var answer string
	for i := range result {
		if result[i].Number == number {
			answer = result[i].Answer
		}
	}

	return answer, nil
}


func getQuiz() ([]Quiz, string, error) {
	response, err := http.Get("http://localhost:8000/quiz")
	if err != nil {
		return nil, "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("unable to get json data: %s", response.Status)
	}

	var result []Quiz
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, "", err
	}

	return result, "", nil
}
