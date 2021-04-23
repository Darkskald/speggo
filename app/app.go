package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
	spec, _ := a.SfgService.GetSfgById(name)
	c.JSON(http.StatusOK, spec)
}

func (a App) ListSfgHandler(c *gin.Context) {
	spectra, err := a.SfgService.ListSfgSpectra()
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, spectra)
}
