package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Renderer = loadTemplate()
	e.Use(middleware.Gzip())
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLength: 32,
		TokenLookup: "form:csrftoken",
		ContextKey: "csrf",
		CookieName: "_csrf",
		CookieMaxAge: 86400,
	}))

	e.Static("/", "src/notes/static")

	execute(e)

	e.Logger.Fatal(e.Start(":5000"))
}
