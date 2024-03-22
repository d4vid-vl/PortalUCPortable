package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func LoginView() fyne.CanvasObject {
	inicio := widget.NewLabel("Bienvenido a Portal UC")
	input_user := widget.NewEntry()
	input_password := widget.NewPasswordEntry()
	button_login := widget.NewButton("Iniciar Sesión", func() {
	})

	input_user.SetPlaceHolder("Ejemplo: \"jmiranda\"")
	input_password.SetPlaceHolder("Ejemplo: \"MiPerrito123\"")

	content := container.NewVBox(inicio, input_user, input_password, button_login)

	return content
}
