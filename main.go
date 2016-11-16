package main

import (
	"github.com/andlabs/ui"
	"github.com/baseigneurie/desktop-sms/dialer"
)

func main() {
	err := ui.Main(func() {
		name := ui.NewEntry()
		box := ui.NewVerticalBox()
		dlr := dialer.GetDialer()

		box.Append(ui.NewLabel("Enter contact name:"), false)
		box.Append(name, false)

		box.Append(dlr, false)
		window := ui.NewWindow("Hello", 300, 200, false)
		window.SetChild(box)

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
