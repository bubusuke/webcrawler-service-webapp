package folio

import (
	"github.com/bubusuke/webcrawler-service-webapp/db"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Themes represents Folio's Investment-Theme List.
type Themes []Theme

// Theme represents one of Folio's Investment-Themes.
type Theme struct {
	// ThemeID is URI resource name of the theme.
	ThemeID string `db:"theme_id"`
	// Title is Theme Name. Title may have multi byte character.
	Title string `db:"title"`
	// IsSelected represent whether it is selected on the web application.
	// It is prepared for html drawing.
	IsSelected bool
}

// ThemeDetail represents the investment stocks in the Folio Investment Themes.
type ThemeDetail struct {
	// Title is the theme's name. Not stock name.
	Title string
	// Stocks represents name of stocks in the Folio Investment Themes.
	Stocks []string
}

// GetThemes gets Themes from DB and returns them.
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

// getTitle returns the theme name that corresponds to URI resource name.
func (ths *Themes) getTitle(queryID string) string {
	for _, th := range *ths {
		if th.ThemeID == queryID {
			return th.Title
		}
	}
	return ""
}

// GetThemesDetails gets name of stocks that corresponds to the theme from DB.
// And returns them.
func (ths *Themes) GetThemesDetails(queryID string) (ThemeDetail, error) {

	if ths.getTitle(queryID) == "" {
		// Initial display corresponds to this case.
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
