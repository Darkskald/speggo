package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"speggo/app/models"
	"speggo/ports"
)

type App struct {
	R          *gin.Engine
	SfgService ports.SfgPort
}

func NewApp(ssv ports.SfgPort) App {

	app := App{
		R:          gin.Default(),
		SfgService: ssv,
	}

	app.R.Use(CORSMiddleware())

	// load HTML templates and static files
	app.R.Static("./assets", "templates/assets")
	app.R.LoadHTMLGlob("templates/*.html")

	return app
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")

		c.Next()
	}
}

func (a App) GetSfgByNameHandler(c *gin.Context) {
	name := c.Param("name")
	spec, err := a.SfgService.GetSfgById(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JsonError{
			Code:    http.StatusInternalServerError,
			Message: "Pech für dich, da lief was schief!",
		})
	} else {
		c.JSON(http.StatusOK, spec)
	}

}

func (a App) ListSfgHandler(c *gin.Context) {
	spectra, err := a.SfgService.ListSfgSpectra()
	// test := c.Request.URL.Query()
	// log.Println(test)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, models.JsonError{
			Code:    http.StatusInternalServerError,
			Message: "Pech für dich, da lief was schief!",
		})
	}

	c.JSON(http.StatusOK, spectra)
}

func (a App) IndexHandler(c *gin.Context) {
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the old_index.html template
		"index.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title": "Home Page",
		},
	)

}

// todo: check if query param is empty string
// todo implement dispatcher for different query parameters
