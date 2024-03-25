package views

import (
	"PortalUCPortable/assets"
	"PortalUCPortable/connections/api"
	"PortalUCPortable/src/functions"

	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
)

func LoginView(w fyne.Window) fyne.CanvasObject {

	inicio := widget.NewLabel("Bienvenido a Portal UC")
	inicio.Alignment = fyne.TextAlignCenter
	inicio.Importance = widget.HighImportance
	inicio.TextStyle = fyne.TextStyle{Bold: true}

	autentificación := canvas.NewImageFromResource(assets.PUCLogin)
	c_autentificación := container.NewGridWrap(fyne.NewSize(400, 125), autentificación)

	// Entries
	input_user := widget.NewEntry()
	input_user.Validator = validation.NewRegexp(`^[A-Za-z0-9_-]+$`, "El usuario solo puede tener letras, números, '-' y '_'")
	input_user.SetPlaceHolder("Ejemplo: \"jmiranda\"")
	input_password := widget.NewPasswordEntry()
	input_password.Validator = validation.NewRegexp(`^[A-Za-z0-9_\-!?#$%&/]+$`, "La contraseña solo puede tener letras, números, y algunos simbolos especiales")
	input_password.SetPlaceHolder("Ejemplo: \"MiPerrito123\"")

	// Labels de alertas
	labelError := canvas.NewText("Credenciales Inválidas", color.NRGBA{R: 149, G: 0, B: 0, A: 255})
	labelError.Hide()
	labelÉxito := canvas.NewText("Cargando...", color.NRGBA{R: 0, G: 149, B: 0, A: 255})
	labelÉxito.Hide()

	// Check para guardar contraseña para otra futura sesión
	checkDatos := widget.NewCheck("Guardar Datos", func(bool) {
		// TODO: Hacer un guardar contraseña versatil y útil
	})

	// Botón para loguearse
	button_login := widget.NewButton("Iniciar Sesión", func() {
		resp := api.LoginPortal(input_user.Text, input_password.Text)
		fmt.Println(resp)
		if resp {
			labelError.Show()
			labelÉxito.Hide()
		} else {
			labelError.Hide()
			labelÉxito.Show()
			time.Sleep(2 * time.Second)
			w.SetContent(functions.LoaderView())
		}
	})
	button_color := canvas.NewRectangle(color.NRGBA{R: 0, G: 149, B: 0, A: 100})
	button_color.CornerRadius = 5

	// Validación para usuario y contraseña
	validateInputs := func() {
		if input_user.Validate() == nil && input_password.Validate() == nil {
			button_login.Enable()
		} else {
			button_login.Disable()
		}
	}

	input_user.OnChanged = func(text string) {
		validateInputs()
	}
	input_password.OnChanged = func(text string) {
		validateInputs()
	}

	button_login.Disable()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Usuario", Widget: input_user},
			{Text: "Contraseña", Widget: input_password},
		},
	}

	content := container.NewVBox(inicio, c_autentificación, form, labelError, labelÉxito, checkDatos, container.NewStack(button_login, button_color))

	return content
}
