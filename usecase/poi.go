package usecase

import (
	"github.com/labstack/echo"
	"github.com/rs/xid"
	"notes/db"
	"notes/models"
	"time"
)

type Header struct {
	Title       string
	Keywords    string
	Description string
}

type PoiData struct {
	HeaderInfo Header
	CSRFToken  string
	Note       *models.Note
	URL        string
	BaseURL    string
	Page       string
}

// Show Index
func GetPoiIndex(c echo.Context) (data PoiData) {
	request := c.Request()
	base := "https://" + request.Host
	path := request.RequestURI

	data = PoiData{
		HeaderInfo: Header{
			Title:       "ポイペ | /tmp/notes",
			Keywords:    "テンプノート,/tmp/notes,一週間,期間限定,メモ,ノート,memo,note,7日間,OSS,オープンソース,ポイ,簡易メモ帳,共有,シェア,WEB,ウェブ,投稿,ポイする,ポイペ",
			Description: "表示期間限定型ウェブメモ帳「/tmp/notes（テンプノート）」のポイペ（メモ帳投稿ページ）です",
		},
		CSRFToken: c.Get("csrf").(string),
		Page:      "poi",
		BaseURL:   base,
		URL:       base + path,
	}
	return
}

// Get Poi detail
func GetPoiDetail(c echo.Context) (data PoiData) {
	id := c.Param("id")
	note := models.NewRepo(db.Connect()).FindByID(id)

	request := c.Request()
	path := "https://" + request.Host + request.RequestURI

	data = PoiData{
		HeaderInfo: Header{
			Title:       "ポイショ | /tmp/notes",
			Keywords:    "テンプノート,/tmp/notes,一週間,期間限定,メモ,ノート,memo,note,7日間,OSS,オープンソース,ポイ,簡易メモ帳,共有,シェア,WEB,ウェブ,投稿,ポイ詳細,ポイショ",
			Description: "表示期間限定型ウェブメモ帳「/tmp/notes（テンプノート）」のポイショ（メモ帳詳細ページ）です",
		},
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
	if title == "" {
		note.Title = nil
	} else {
		note.Title = &title
	}

	note.Content = requests.FormValue("content")
	note.IP = &requests.RemoteAddr
	note.CreatedAt = time.Now()
	note.ExpireAt = time.Now().AddDate(0, 0, 7)

	if note.Content == "" {
		return nil
	}

	con := db.Connect()
	if models.NewRepo(con).PostNote(&note) {
		return &note
	}
	return nil
}
