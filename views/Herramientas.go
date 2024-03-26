package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func HerramientasView() fyne.CanvasObject {
	h := widget.NewLabel("Herramientas")

	container := container.NewVBox(h)

	return container
}
