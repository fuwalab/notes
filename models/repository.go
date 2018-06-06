package models

import (
	"gopkg.in/mgo.v2"
	"fmt"
)

type Repo struct {
	db *mgo.Database

}

func NewRepo(con *mgo.Database) *Repo {
	return &Repo{db: con}
}

func (r *Repo) GetAllNotes() (results []*Note) {
	err := r.db.C("notes").Find(nil).All(&results)
	if err != nil {
		fmt.Printf("log: %v", err)
	}
	return
}
