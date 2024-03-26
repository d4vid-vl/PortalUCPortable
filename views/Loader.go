package views

import (
	"fyne.io/fyne/v2/container"
)

func Loader() *container.AppTabs {
	tabs := container.NewAppTabs(
		container.NewTabItem("Inicio", InicioView()),
		container.NewTabItem("Datos Personales", DatosView()),
		container.NewTabItem("Información Académica", InfoAcademicaView()),
		container.NewTabItem("Información Financiera", InfoFinancieraView()),
		container.NewTabItem("Herramientas", HerramientasView()),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	return tabs
}
