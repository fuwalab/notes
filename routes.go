package main

import (
	"github.com/labstack/echo"
	"net/http"
	"notes/usecase"
)

func execute(e *echo.Echo) {
	// index
	e.GET("/", func(c echo.Context) error {
		data := struct {
			Page string
		}{
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
		if data.Note == nil {
			return c.Render(http.StatusNotFound, "404", nil)
		}
		return c.Render(http.StatusOK, "layout", data)
	})

	// Post Poi data
	e.POST("/poi/", func(c echo.Context) error {
		var id string
		note := usecase.PostPoiDetail(c)
		if note != nil {
			id = note.ID
		}
		return c.Redirect(http.StatusMovedPermanently, "/poi/"+id)
	})
}
