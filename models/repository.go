package models

import (
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Repo struct {
	db *mgo.Database
}

func NewRepo(con *mgo.Database) *Repo {
	return &Repo{db: con}
}

// Retrieve All notes data
func (r *Repo) GetAllNotes() (results []*Note) {
	err := r.db.C("notes").Find(nil).All(&results)
	if err != nil {
		panic(err.Error())
	}
	return
}

// Retrieve note
func (r *Repo) FindByID(id string) (ret *Note) {
	err := r.db.C("notes").Find(bson.M{
		"ID": bson.M{
			"$eq": id,
		},
		"ExpireAt": bson.M{
			"$gte": time.Now(),
		},
	}).One(&ret)

	if err != nil {
		log.Error(err.Error())
		return nil
	}
	return
}

// Save note data
func (r *Repo) PostNote(data *Note) bool {
	err := r.db.C("notes").Insert(data)
	if err != nil {
		panic(err.Error())
	}
	return true
}
