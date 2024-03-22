package functions

import (
	"fyne.io/fyne/v2"
)

func Startup(a fyne.App, widget fyne.CanvasObject) {
	w := a.NewWindow("Portal UC")

	w.Resize(fyne.NewSize(600, 600))
	w.SetFixedSize(true)

	w.SetContent(widget)
	w.ShowAndRun()
}
