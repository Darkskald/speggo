package main

import (
	app2 "speggo/app"
	"speggo/services/sqlite"
)

func main() {
	dbPath := "/home/flo/coding/goprojects/speggo/spectra.db"

	sfgService := sqlite.NewSfgService(dbPath)
	app := app2.NewApp(&sfgService)
	app.SetupRoutes()

	app.R.Run(":8080")
}
