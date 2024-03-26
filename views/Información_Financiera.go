package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func InfoFinancieraView() fyne.CanvasObject {
	h := widget.NewLabel("Información Financiera")

	container := container.NewVBox(h)

	return container
}
