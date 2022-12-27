package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	"strconv"
)

func main() {
	a := app.New()
	l := widget.NewLabel("Hello Fyne!")
	e := widget.NewEntry()
	e.SetText("0")
	w := a.NewWindow("Hello")
	//w.SetContent(
	//	container.NewVBox(
	//		widget.NewLabel("Hello Fyne!"),
	//		widget.NewLabel("This is sampleapplication!"),
	//	),
	//)
	w.SetContent(
		container.NewVBox(
			l, e,
			widget.NewButton("Click Me!", func() {
				n, _ := strconv.Atoi(e.Text)
				l.SetText("Total: " + strconv.Itoa(total(n)))
			}),
		),
	)

	w.ShowAndRun()
}

func total(n int) int {
	t := 0
	for i := 0; i <= n; i++ {
		t += i
	}
	return t
}
