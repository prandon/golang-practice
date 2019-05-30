package dao

import (
	//"fmt"
	"log"
	model "../model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MoviesDao struct {
	Server string
	Database string
}

const (
	COLLECTION = "movies"
)

var db *mgo.Database

func (m *MoviesDao) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *MoviesDao) FindAll() ([]model.Movie, error) {
	var movies []model.Movie
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}