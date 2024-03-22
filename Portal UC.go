package main

import (
	"PortalUCPortable/src/functions"
	"PortalUCPortable/views"

	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Portal UC")
	functions.Startup(a, w, views.LoginView(w))
}
