package mongo

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"morses-code/mc-quizbot/quiz"
)

type Mongo struct {
	Collection string
	Database   string
	Session    *mgo.Session
	URI        string
}

func (m *Mongo) Init() (session *mgo.Session, err error) {
	if session != nil {
		return nil, errors.New("session already exists")
	}

	if session, err = mgo.Dial(m.URI); err != nil {
		return nil, err
	}

	session.EnsureSafe(&mgo.Safe{WMode: "majority"})
	session.SetMode(mgo.Strong, true)
	return session, nil
}

func (m *Mongo) GetQuestions() ([]quiz.Question, error) {
	session := m.Session.Copy()
	defer session.Close()

	var result []quiz.Question
	c := session.DB("quiz").C("questions")
	err := c.Find(nil).All(&result)
	if err != nil {
		log.Panic(err)
	}

	return result, nil
}

func (m *Mongo) GetQuestion(number int) (string, error) {
	session := m.Session.Copy()
	defer session.Close()

	result := quiz.Question{}
	c := session.DB("quiz").C("questions")
	err := c.Find(bson.M{"number": number}).One(&result)
	if err != nil {
		log.Panic(err)
	}

	return result.Question, nil
}