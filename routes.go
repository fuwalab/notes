package main

import (
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type User struct {
	Name string
	Note string
	IP   *string
}

func execute(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		session, _ := mgo.Dial("mongo")
		db := session.DB("drafts")

		var results []User
		db.C("users").Find(bson.M{}).All(&results)

		for _, c := range results {
			fmt.Printf("retrieve from mongodb: %v\n", c)
		}

		data := struct {
			Contents string
			Page string
		} {
			Contents: "ページのコンテンツ",
			Page: "index",
		}
		return c.Render(http.StatusOK, "layout", data)
	})

	e.GET("/poi/", func(c echo.Context) error {
		data := struct {
			Contents string
			Page string
		} {
			Contents: "ポイする",
			Page: "poi",
		}
		return c.Render(http.StatusOK, "layout", data)
	})
}
