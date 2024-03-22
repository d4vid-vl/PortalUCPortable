package views

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func LoaderView() *container.AppTabs {
	tabs := container.NewAppTabs(
		container.NewTabItem("Inicio", LoginView()),
		container.NewTabItem("Datos Personales", widget.NewLabel("Datos Personales")),
		container.NewTabItem("Información Académica", widget.NewLabel("Información Académica")),
		container.NewTabItem("Información Financiera", widget.NewLabel("Información Financiera")),
		container.NewTabItem("Herramientas", widget.NewLabel("Herramientas")),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	return tabs
}
