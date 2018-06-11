package main

import (
	"github.com/labstack/echo"
	"net/http"
	"notes/usecase"
)

func execute(e *echo.Echo) {
	// index
	e.GET("/", func(c echo.Context) error {
		request := c.Request()
		base := "https://" + request.Host
		path := request.RequestURI

		data := usecase.PoiData{
			HeaderInfo: usecase.Header{
				Title:       "/tmp/notes | ウェブに残せるちょっとしたメモ帳",
				Keywords:    "テンプノート,/tmp/notes,一週間,期間限定,メモ,ノート,memo,note,7日間,OSS,オープンソース,ポイ,簡易メモ帳,共有,シェア",
				Description: "/tmp/notes(テンプノート)はしっかり記録するほどでもない「ちょっとだけメモを取っておきたい」をカタチにする、ウェブ投稿型のメモ帳サービスです。メモの閲覧可能期間は1週間。投稿から1週間が経過すると自動的に表示されなくなります。",
			},
			Page:    "index",
			BaseURL: base,
			URL:     base + path,
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
		if data.Notes == nil {
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

	// tag list
	e.GET("/tag/", func(c echo.Context) error {
		data := usecase.GetTagIndex(c)
		return c.Render(http.StatusOK, "layout", data)
	})

	//
	e.POST("/tag/", func(c echo.Context) error {
		tag := usecase.PostTagDetail(c)
		return c.Redirect(http.StatusMovedPermanently, "/tag/"+tag)
	})

	// tag detail
	e.GET("/tag/:id", func(c echo.Context) error {
		data := usecase.GetTagDetail(c)

		return c.Render(http.StatusOK, "layout", data)
	})
}
