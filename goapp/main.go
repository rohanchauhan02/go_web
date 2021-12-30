package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("My Name is Rohan")
	myWindow.SetContent(widget.NewLabel("I am not a terrorist"))
	myWindow.Resize(fyne.NewSize(350, 500))
	

	go showAnother(myApp)
	myWindow.ShowAndRun()
}

func showAnother(a fyne.App) {
	time.Sleep(time.Second * 3)

	win := a.NewWindow("BUT")
	win.SetContent(widget.NewLabel("I know destruction"))
	win.Resize(fyne.NewSize(250, 250))
	win.Show()

	time.Sleep(time.Second * 5)
	win.Close()
}