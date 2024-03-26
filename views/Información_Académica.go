package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func InfoAcademicaView() fyne.CanvasObject {
	h := widget.NewLabel("Información Académica")

	container := container.NewVBox(h)

	return container
}
