package terminalview

import (
	"lib/model"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func GetSerialConfig() *model.SerialConfig {
	app := tview.NewApplication()
	config := &model.SerialConfig{}

	form := tview.NewForm()
	form.SetBorder(true).
		SetTitle("DialIt").
		SetTitleAlign(tview.AlignCenter)

	// Add COM Port input field
	form.AddInputField("COM Port", "", 20, nil, func(text string) {
		config.ComPort = text
	})

	// Add Baud Rate input field
	form.AddInputField("Baud Rate", "", 20, nil, func(text string) {
		br, _ := strconv.ParseUint(text, 10, 32)
		config.BaudRate = uint(br)
	})

	// Add buttons
	form.AddButton("Save", func() {
		app.Stop()
	})
	form.AddButton("Quit", func() {
		app.Stop()
	})

	// Set form appearance
	form.SetButtonsAlign(tview.AlignCenter)
	form.SetFieldBackgroundColor(tcell.ColorDefault)
	form.SetButtonBackgroundColor(tcell.ColorBlue)

	// Center the form
	flex := tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().
			AddItem(nil, 0, 1, false).
			AddItem(form, 40, 1, true).
			AddItem(nil, 0, 1, false), 10, 1, true).
		AddItem(nil, 0, 1, false)

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	return config
}
