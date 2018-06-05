package usecase

import (
	"notes/models"
	"github.com/labstack/echo"
)

// Show Index
func GetPoiIndex() (data interface{}) {
	data = struct {
		Contents string
		Page string
	} {
		Contents: "ポイする",
		Page: "poi",
	}
	return
}

type PoiData struct {
	Contents string
	Note models.Note
	Page string
}

// Poi detail
func GetPoiDetail(c echo.Context) (data PoiData) {
	//params := c.Param("id")
	//requests := c.Request()
	//fmt.Printf("params: %v\n", params)
	//fmt.Printf("requests: %v\n", requests)

	title := "サンプルだよ"
	note := models.Note{
		ObjectID: "xxx",
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
