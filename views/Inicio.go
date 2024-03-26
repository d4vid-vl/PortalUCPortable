package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func InicioView() fyne.CanvasObject {
	h := widget.NewLabel("Inicio")

	container := container.NewVBox(h)

	return container
}
