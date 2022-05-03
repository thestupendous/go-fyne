package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	a "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fmt.Println("vim-go")

	app := a.New()
	window := app.NewWindow("This is Title")
	window.Resize(fyne.NewSize(800, 600))

	//adding text label
	titleLabel := widget.NewLabel("My App Which is Cool")
	window.SetContent(titleLabel)

	//adding button
	button1 := widget.NewButton("Press Me", func() {
		fmt.Println("button1 was pressed")
	})
	window.SetContent(button1)

	//adding checkbox
	cb1 := widget.NewCheck("Are you fine", func(pressed bool) {
		if pressed {
			fmt.Println("user is OK, exiting")
		} else {
			fmt.Println("user is not OK, call the POLICE!!")
		}
	})
	window.SetContent(cb1)

	window.ShowAndRun()
}
