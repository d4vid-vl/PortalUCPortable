package functions

import (
	"fyne.io/fyne/v2"
)

func Startup(a fyne.App, w fyne.Window, widget fyne.CanvasObject) {
	w.Resize(fyne.NewSize(600, 600))
	w.SetFixedSize(true)

	w.SetContent(widget)
	w.ShowAndRun()
}
