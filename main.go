package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/bubusuke/webcrawler-service-webapp/folio"
	"github.com/gin-gonic/gin"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

var appPort string

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.SetPrefix("webapp ")
	appPort = os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	// Either A or B can be accepted.
	appPort = strings.Replace(appPort, ":", "", 1)
	appPort = ":" + appPort
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")

	//URI handle definition.
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})
	r.GET("/folio", func(c *gin.Context) {
		// Parse query parameter.
		qThemeID := c.Query("theme")

		// Getting response parameters.
		ths, err := folio.GetThemes(qThemeID)
		if err != nil {
			log.Println(err)
		}
		td, err := ths.GetThemesDetails(qThemeID)
		if err != nil {
			log.Println(err)
		}

		// Response format definition.
		c.HTML(http.StatusOK, "folio.tmpl", gin.H{
			"themes":      ths,
			"themeDetail": td,
		})
	})

	// Start to run.
	r.Run(appPort)
}
