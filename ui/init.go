package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func Initialize(app fyne.App) (fyne.Window, *AppWindow) {

	appText := "Shadow of Text"
	appComponents := NewAppWindow(appText)
	window := app.NewWindow(appText)

	mainContainer := container.NewVBox(
		appComponents.Menu,
		appComponents.InputBox,
		appComponents.TextGrid,
	)

	window.SetContent(mainContainer)
	window.Resize(fyne.NewSize(800, 600))

	return window, appComponents
}
