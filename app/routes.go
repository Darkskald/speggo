package app

func (a *App) SetupRoutes() {
	a.R.GET("/", a.IndexHandler)
	a.R.GET("/sfg/:name", a.GetSfgByNameHandler)
	a.R.POST("/sfg", a.PostSfgHandler)
	a.R.GET("/sfg", a.ListSfgHandler)
	a.R.GET("/sfg/fuzzy/:filter", a.SearchHandler)
}
