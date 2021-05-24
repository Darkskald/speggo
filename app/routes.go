package app

func (a *App) SetupRoutes() {
	a.R.GET("/", a.IndexHandler)
	a.R.GET("/sfg/:name", a.GetSfgByNameHandler)
	a.R.GET("/sfg", a.ListSfgHandler)
}
