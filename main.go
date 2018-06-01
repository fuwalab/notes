package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Renderer = loadTemplate()
	e.Use(middleware.Gzip())

	e.Static("/", "src/notes/static")

	execute(e)

	e.Logger.Fatal(e.Start(":5000"))
}
