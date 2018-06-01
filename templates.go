package main

import (
	"io"
	"html/template"
	"github.com/labstack/echo"
)

type Template struct {
	Template *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Template.ExecuteTemplate(w, name, data)
}

func loadTemplate() *Template{
	t := &Template{
		Template: template.Must(template.ParseGlob("src/notes/templates/*.html")),
	}
	return t
}
