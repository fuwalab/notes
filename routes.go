package main

import (
	"github.com/labstack/echo"
	"fmt"
	"net/http"
	"notes/db"
	"notes/models"
)

func execute(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		con := db.Connect()

		results := models.NewRepo(con).GetAllNotes()

		for _, c := range results {
			fmt.Printf("%v\n", c.ID)
			fmt.Printf("%v\n", *c.Title)
			fmt.Printf("%v\n", c.Content)
			fmt.Printf("%v\n", c.CreatedAt)
			fmt.Printf("%v\n", *c.UA)
			fmt.Printf("%v\n", *c.IP)
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
