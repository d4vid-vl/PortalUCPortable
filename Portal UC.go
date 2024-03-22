package main

import (
	"PortalUCPortable/src/functions"
	"PortalUCPortable/views"

	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	functions.Startup(a, views.LoaderView())
}
