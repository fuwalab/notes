package db

import (
	"gopkg.in/mgo.v2"
	"notes/config"
	"time"
)

func Connect() (db *mgo.Database) {
	mongoInfo := &mgo.DialInfo{
		Addrs:    []string{"mongo"},
		Timeout:  60 * time.Second,
		Database: config.DBName,
		Username: config.DBUser,
		Password: config.DBPassword,
		Source:   config.DBSource,
	}
	session, err := mgo.DialWithInfo(mongoInfo)
	if err != nil {
		panic(err)
	}

	db = session.DB("")
	return
}
