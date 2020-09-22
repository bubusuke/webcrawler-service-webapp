package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/bubusuke/webcrawler-service-webapp/folio"
	"github.com/gin-gonic/gin"
)

// appPort is port parameter of this web application.
// default value is :8080.
var appPort string

// init sets log format and read env.
// APP_PORT accepts either ":XXXX" or "XXXX".
func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.SetPrefix("webapp ")
	appPort = os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	// Either ":XXXX" or "XXXX" can be accepted.
	appPort = strings.Replace(appPort, ":", "", 1)
	appPort = ":" + appPort
}

// main starts server process and handles request.
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
