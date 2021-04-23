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
	return App{
		R:          gin.Default(),
		SfgService: ssv,
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

// todo: check if query param is empty string
// todo implement dispatcher for different query parameters
