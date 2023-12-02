package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type AppWindow struct {
	Menu     *fyne.Container
	InputBox *widget.Entry
	TextGrid *widget.TextGrid
}

func NewAppWindow(input string) *AppWindow {
	return &AppWindow{
		Menu:     buildMenu(),
		InputBox: buildInputBox(),
		TextGrid: buildTextGrid(input),
	}
}

func buildMenu() *fyne.Container {
	return container.NewWithoutLayout()
}

func buildInputBox() *widget.Entry {
	return widget.NewMultiLineEntry()
}

func buildTextGrid(input string) *widget.TextGrid {

	return widget.NewTextGridFromString(input)
}
