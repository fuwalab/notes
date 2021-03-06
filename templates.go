package main

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
	"strings"
	"time"
)

type Template struct {
	Template *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Template.ExecuteTemplate(w, name, data)
}

func loadTemplate() *Template {
	templateFunctions := template.FuncMap{
		"nl2br": nl2br,
		"date":  date,
	}

	t := &Template{
		Template: template.Must(template.New("").Funcs(templateFunctions).ParseGlob("src/notes/templates/*.html")),
	}
	return t
}

// replace "\n" to "<br>"
func nl2br(text string) template.HTML {
	return template.HTML(strings.Replace(template.HTMLEscapeString(text), "\n", "<br>", -1))
}

// format date
func date(dt time.Time) string {
	location := time.FixedZone("Asia/Tokyo", 9*60*60)
	return dt.In(location).Format("2006/01/02 15:04")
}
