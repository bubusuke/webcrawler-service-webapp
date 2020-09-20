package main

import (
	"html/template"
	"log"
	"net/http"
	"sync"

	"github.com/bubusuke/webcrawler-service/folio"
	"github.com/gin-gonic/gin"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})
	r.GET("/folio", func(c *gin.Context) {
		qTheme := c.Query("theme")
		ths := folio.ReadThemes(qTheme, "./folio/themes.html")
		td, err := ths.CrawlThemesDetail(qTheme)
		if err != nil {
			log.Println(err)
		}

		c.HTML(http.StatusOK, "folio.tmpl", gin.H{
			"themes":      ths,
			"themeDetail": td,
		})
	})

	r.Run(":8080")
}
