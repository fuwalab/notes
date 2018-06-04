package db

import (
	"gopkg.in/mgo.v2"
	"time"
)

func Connect() (db *mgo.Database) {
	mongoInfo := &mgo.DialInfo{
		Addrs:    []string{"mongo"},
		Timeout:  60 * time.Second,
		Database: "",
		Username: "",
		Password: "",
		Source:   "",
	}
	session, err := mgo.DialWithInfo(mongoInfo)
	if err != nil {
		panic(err)
	}

	db = session.DB("")
	return
}
