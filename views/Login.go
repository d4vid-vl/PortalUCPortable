package views

import (
	"PortalUCPortable/connections/api"
	"PortalUCPortable/src/functions"
	"fmt"

	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func LoginView(w fyne.Window) fyne.CanvasObject {

	inicio := widget.NewLabel("Bienvenido a Portal UC")
	input_user := widget.NewEntry()
	input_password := widget.NewPasswordEntry()

	labelError := canvas.NewText("Credenciales Inválidas", color.NRGBA{R: 149, G: 0, B: 0, A: 255})
	labelError.Hide()
	labelÉxito := canvas.NewText("Cargando...", color.NRGBA{R: 0, G: 149, B: 0, A: 255})
	labelÉxito.Hide()

	button_login := widget.NewButton("Iniciar Sesión", func() {
		resp := api.LoginPortal(input_user.Text, input_password.Text)
		fmt.Println(resp)
		if resp {
			labelError.Show()
			labelÉxito.Hide()
		} else {
			labelError.Hide()
			labelÉxito.Show()
			w.SetContent(functions.LoaderView())
		}
	})

	input_user.SetPlaceHolder("Ejemplo: \"jmiranda\"")
	input_password.SetPlaceHolder("Ejemplo: \"MiPerrito123\"")

	content := container.NewVBox(inicio, input_user, input_password, labelError, labelÉxito, button_login)

	return content
}
