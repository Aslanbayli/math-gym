package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app = tview.NewApplication()
var startText = tview.NewTextView().SetTextColor(tcell.ColorGreen).SetText("(t) to start training\n(q) to quit")
var box = tview.NewBox().SetBorder(true).SetTitle(" ----- 💪 Welcome to the Math GYM 💪 ----- ")
var flex = tview.NewFlex()
var pages = tview.NewPages()
var flexContainer = tview.NewFlex()

func main() {

	dummy := tview.NewTextView().
		SetTextColor(tcell.ColorRed).
		SetText("Hello!")

	flex.SetDirection(tview.FlexRow).
		AddItem(dummy, 0, 1, true)

	box.SetDrawFunc(func(screen tcell.Screen, x, y, w, h int) (int, int, int, int) {
		//tview.Print(screen, dummy.GetText(true), x, y+10, w, tview.AlignCenter, tcell.ColorRed)
		flex.Draw(screen)
		return x, y, w, h
	})

	flexContainer.SetDirection(tview.FlexRow).
		AddItem(box, 0, 1, true)

	pages.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			app.Stop()
		} else if event.Rune() == 't' {
			pages.SwitchToPage("Main")
		}
		return event
	})

	pages.AddPage("Start Menu", startText, true, true).
		AddPage("Main", flexContainer, true, false)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
