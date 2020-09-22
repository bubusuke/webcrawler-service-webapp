package folio

import (
	"github.com/bubusuke/webcrawler-service-webapp/db"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Themes []Theme

type Theme struct {
	ThemeID    string `db:"theme_id"`
	Title      string `db:"title"`
	IsSelected bool
}

type ThemeDetail struct {
	Title  string
	Stocks []string
}

func GetThemes(queryID string) (Themes, error) {

	query := "SELECT theme_id, title FROM themes ORDER BY seq"
	db, err := sqlx.Connect("postgres", db.GetDbInfo())
	if err != nil {
		return nil, err
	}
	rows, err := db.Queryx(query)
	if err != nil {
		return nil, err
	}

	ths := make([]Theme, 0, 100)
	th := Theme{}
	for rows.Next() {
		err := rows.StructScan(&th)
		if err != nil {
			return nil, err
		}
		if th.ThemeID == queryID {
			th.IsSelected = true
		} else {
			th.IsSelected = false
		}
		ths = append(ths, th)
	}
	return ths, nil
}

func (ths *Themes) getTitle(queryID string) string {
	for _, th := range *ths {
		if th.ThemeID == queryID {
			return th.Title
		}
	}
	return ""
}

func (ths *Themes) GetThemesDetails(queryID string) (ThemeDetail, error) {
	if ths.getTitle(queryID) == "" {
		return ThemeDetail{}, nil
	}

	td := ThemeDetail{
		Title:  ths.getTitle(queryID),
		Stocks: make([]string, 0, 20),
	}
	query := "SELECT title FROM theme_details WHERE theme_id = :theme_id ORDER BY detail_id"
	db, err := sqlx.Connect("postgres", db.GetDbInfo())
	if err != nil {
		return ThemeDetail{}, err
	}
	rows, err := db.NamedQuery(query, map[string]interface{}{"theme_id": queryID})
	if err != nil {
		return ThemeDetail{}, err
	}
	var stock string
	for rows.Next() {
		if err := rows.Scan(&stock); err != nil {
			return ThemeDetail{}, err
		}
		td.Stocks = append(td.Stocks, stock)
	}
	return td, nil
}
