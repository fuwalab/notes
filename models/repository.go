package models

import (
	"gopkg.in/mgo.v2"
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

// Save note data
func (r *Repo) PostNote(data *Note) bool {
	err := r.db.C("notes").Insert(data)
	if err != nil {
		panic(err.Error())
	}
	return true
}
