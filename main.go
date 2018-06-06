package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	e := echo.New()
	e.Renderer = loadTemplate()
	e.Use(middleware.Gzip())
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLength:  32,
		TokenLookup:  "form:csrftoken",
		ContextKey:   "csrf",
		CookieName:   "_csrf",
		CookieMaxAge: 30 * 60,
	}))

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			if he.Code == 404 {
				c.Render(http.StatusNotFound, "404", nil)
			} else {
				c.Render(http.StatusInternalServerError, "500", nil)
			}
		}
	}

	e.Static("/", "src/notes/static")
	e.File("/favicon.ico", "src/notes/static/images/favicon.ico")

	execute(e)

	e.Logger.Fatal(e.Start(":5000"))
}
