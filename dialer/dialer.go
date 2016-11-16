package dialer

import (
	"github.com/andlabs/ui"
)

func GetDialer() ui.Box {
	dialer := ui.NewVerticalBox()
	d := loadDialer(*dialer)
	return d
}

func loadDialer(dialer ui.Box) ui.Box {
	topRow := ui.NewHorizontalBox()
	midRow := ui.NewHorizontalBox()
	botRow := ui.NewHorizontalBox()
	zeroRow := ui.NewHorizontalBox()
	// fullBox := ui.NewVerticalBox()

	bt1 := ui.NewButton("1")
	bt2 := ui.NewButton("2")
	bt3 := ui.NewButton("3")
	bt4 := ui.NewButton("4")
	bt5 := ui.NewButton("5")
	bt6 := ui.NewButton("6")
	bt7 := ui.NewButton("7")
	bt8 := ui.NewButton("8")
	bt9 := ui.NewButton("9")
	bt0 := ui.NewButton("0")

	// Top Row
	topRow.Append(bt1, true)
	topRow.Append(bt2, true)
	topRow.Append(bt3, true)

	// Middle Row
	midRow.Append(bt4, true)
	midRow.Append(bt5, true)
	midRow.Append(bt6, true)

	// Bottom Row
	botRow.Append(bt7, true)
	botRow.Append(bt8, true)
	botRow.Append(bt9, true)

	//Zero Row
	zeroRow.Append(bt0, true)

	// Buttons container
	dialer.Append(topRow, true)
	dialer.Append(midRow, true)
	dialer.Append(botRow, true)
	dialer.Append(zeroRow, true)

	return dialer
}
