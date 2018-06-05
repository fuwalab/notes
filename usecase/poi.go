package usecase

import (
	"notes/models"
	"github.com/labstack/echo"
	"notes/db"
	"time"
	"fmt"
)

// Show Index
func GetPoiIndex(c echo.Context) (data interface{}) {
	fmt.Printf("%v\n",  c.Get("csrf"))
	data = struct {
		Contents string
		CSRFToken string
		Page string
	} {
		Contents: "ポイする",
		CSRFToken: c.Get("csrf").(string),
		Page: "poi",
	}
	return
}

type PoiData struct {
	Contents string
	Note models.Note
	Page string
}

// Get Poi detail
func GetPoiDetail(c echo.Context) (data PoiData) {
	//params := c.Param("id")
	//requests := c.Request()
	//fmt.Printf("params: %v\n", params)
	//fmt.Printf("requests: %v\n", requests)

	title := "サンプルだよ"
	note := models.Note{
		ID: "xxx",
		Title: &title,
		Content: "hogeっていうかこれがコンテンツだよ。",
	}

	data = PoiData {
		Contents: "ポイする",
		Note: note,
		Page: "poi_detail",
	}
	return
}

// Post Poi detail
func PostPoiDetail(c echo.Context) *models.Note {
	// avoid duplicated entry by reloading
	note := models.Note{}
	requests := c.Request()

	ua := requests.UserAgent()
	note.UA = &ua

	title := requests.FormValue("title")
	note.Title = &title

	note.Content = requests.FormValue("content")
	note.IP = &requests.RemoteAddr
	note.CreatedAt = time.Now()

	con := db.Connect()
	if models.NewRepo(con).PostNote(&note) {
		//c.SetCookie(&http.Cookie{Name: "_csrf", Value: ""})
		return &note
	}
	return nil
}
