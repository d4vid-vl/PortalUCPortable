package main

import (
	"AndroidUC/src/functions"
	"AndroidUC/views"

	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	functions.Startup(a, views.LoaderView())
}
