package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func DatosView() fyne.CanvasObject {
	h := widget.NewLabel("Datos Personales")

	container := container.NewVBox(h)

	return container
}
