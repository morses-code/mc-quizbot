package quiz

type Question struct {
	Number   int    `json:"number"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

/*
func GetQuizHandler(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	var result []Question
	c := session.DB("quiz").C("questions")
	err = c.Find(nil).All(&result)
	if err != nil {
		log.Panic(err)
	}


	data, err := json.Marshal(&result)
	if err != nil {
		log.Panic(err)
	}

	fmt.Fprintf(w, string(data))
}*/

/*
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

func CreateQuestionHandler(w http.ResponseWriter, r *http.Request) {
	var question Question
	err := json.NewDecoder(r.Body).Decode(&question)
	if err != nil {
		log.Fatal(err)
	}

	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("quiz").C("questions")
	err = c.Insert(&question)
	if err != nil {
		log.Fatal(err)
	}

	questionNumber := &question.Number

	result := Question{}
	err = c.Find(bson.M{"number": questionNumber}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "successfully created quiz question number: %d", result.Number)
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
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	result := Question{}
	c := session.DB("quiz").C("questions")
	err = c.Find(bson.M{"number": number}).One(&result)
	if err != nil {
		log.Panic(err)
	}

	return result.Question, nil
}

func getAnswer(number int) (string, error) {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	result := Question{}
	c := session.DB("quiz").C("questions")
	err = c.Find(bson.M{"number": number}).One(&result)
	if err != nil {
		log.Panic(err)
	}

	return result.Answer, nil
}
*/
