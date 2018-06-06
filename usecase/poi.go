package usecase

import (
	"github.com/labstack/echo"
	"github.com/rs/xid"
	"notes/db"
	"notes/models"
	"time"
)

// Show Index
func GetPoiIndex(c echo.Context) (data interface{}) {
	data = struct {
		CSRFToken string
		Page      string
	}{
		CSRFToken: c.Get("csrf").(string),
		Page:      "poi",
	}
	return
}

type PoiData struct {
	Note *models.Note
	URL  string
	Page string
}

// Get Poi detail
func GetPoiDetail(c echo.Context) (data PoiData) {
	id := c.Param("id")
	note := models.NewRepo(db.Connect()).FindByID(id)

	request := c.Request()
	path := "://" + request.Host + request.RequestURI

	data = PoiData{
		Note: note,
		URL:  path,
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

	idObject := xid.New()
	note.ID = idObject.String()

	title := requests.FormValue("title")
	note.Title = &title

	note.Content = requests.FormValue("content")
	note.IP = &requests.RemoteAddr
	note.CreatedAt = time.Now()
	note.ExpireAt = time.Now().AddDate(0, 0, 7)

	con := db.Connect()
	if models.NewRepo(con).PostNote(&note) {
		return &note
	}
	return nil
}
