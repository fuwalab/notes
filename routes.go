package main

import (
	"github.com/labstack/echo"
	"fmt"
	"net/http"
	"notes/db"
	"notes/models"
	"notes/usecase"
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

	// Show Poi index page.
	e.GET("/poi/", func(c echo.Context) error {
		data := usecase.GetPoiIndex(c)
		return c.Render(http.StatusOK, "layout", data)
	})

	// Show Poi detail page
	e.GET("/poi/:id", func(c echo.Context) error {
		data := usecase.GetPoiDetail(c)
		return c.Render(http.StatusOK, "layout", data)
	})

	// Post Poi data
	e.POST("/poi/", func(c echo.Context) error {
		note := usecase.PostPoiDetail(c)
		data := struct {
			Contents string
			Note *models.Note
			Page string
		} {
			Contents: "ページのコンテンツ",
			Note: note,
			Page: "poi_detail",
		}

		return c.Render(http.StatusOK, "layout", data)
	})
}
