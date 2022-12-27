package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	"strconv"
)

func main() {
	c := 0
	a := app.New()
	l := widget.NewLabel("Hello Fyne!")
	w := a.NewWindow("Hello")
	//w.SetContent(
	//	container.NewVBox(
	//		widget.NewLabel("Hello Fyne!"),
	//		widget.NewLabel("This is sampleapplication!"),
	//	),
	//)
	w.SetContent(
		container.NewVBox(
			l,
			widget.NewButton("Click Me!", func() {
				c++
				l.SetText("Hello Fyne! " + strconv.Itoa(c))
			}),
		),
	)

	w.ShowAndRun()
}
