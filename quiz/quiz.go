package quiz

type Question struct {
	Number   int    `json:"number"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
