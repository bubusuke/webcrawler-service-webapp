package folio

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Themes []Theme

type Theme struct {
	Id         string
	Title      string
	IsSelected bool
}

type ThemeDetail struct {
	Title  string
	Stocks []string
}

const theme_url = "https://folio-sec.com/theme"

func ReadThemes(queryId string, path string) Themes {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	stringReader := strings.NewReader(string(f))
	doc, err := goquery.NewDocumentFromReader(stringReader)
	if err != nil {
		log.Fatalln(err)
	}
	selection := doc.Find("a.gtm-theme-detail")

	ths := make([]Theme, 0, 100)
	selection.Each(func(i int, s *goquery.Selection) {
		attr, _ := s.Attr("href")
		th := Theme{
			Id:         strings.Replace(attr, "/theme/", "", 1),
			Title:      s.Find("h1").Text(),
			IsSelected: false,
		}
		if th.Id == queryId {
			th.IsSelected = true
		}
		ths = append(ths, th)
	})

	return ths
}

func (ths *Themes) getTitle(queryId string) string {
	for _, th := range *ths {
		if th.Id == queryId {
			return th.Title
		}
	}
	return ""
}

// TODO: 動的サイトのクローラ機能実装が必要
// func CrawlThemes(queryTheme string) []Theme {
// }

func (ths *Themes) CrawlThemesDetail(queryId string) (ThemeDetail, error) {

	td := ThemeDetail{
		Title:  ths.getTitle(queryId),
		Stocks: make([]string, 0, 10),
	}
	if td.Title == "" {
		return ThemeDetail{}, nil
	}

	targetUrl := theme_url + "/" + queryId
	doc, err := goquery.NewDocument(targetUrl)
	if err != nil {
		return ThemeDetail{}, err
	}
	selection := doc.Find("table")
	selection = selection.Find("button.gtm-stock-detail")
	selection.Each(func(i int, s *goquery.Selection) {
		td.Stocks = append(td.Stocks, s.Text())
	})

	return td, nil
}
