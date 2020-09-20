package folio

import (
	"fmt"
	"testing"
)

func Test_ReadThemes(t *testing.T) {
	fmt.Println(ReadThemes("pay", "./themes.html"))
}

func Test_CrawlThemesDetail(t *testing.T) {
	ths := ReadThemes("pay", "./themes.html")
	fmt.Println(ths.CrawlThemesDetail("pay"))
}
