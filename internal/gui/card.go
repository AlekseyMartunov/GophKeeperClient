package gui

import "github.com/rivo/tview"

func newCardPage() *tview.Flex {
	flex := tview.NewFlex()
	form := tview.NewForm()
	infoBox := tview.NewTextArea()

	infoBox.SetText("Add credit card ...", false)
	infoBox.SetBorderPadding(2, 2, 2, 2)

	form.AddInputField("Name", "", 20, nil, nil)
	form.AddInputField("Number", "", 20, nil, nil)
	form.AddInputField("Date", "", 20, nil, nil)
	form.AddInputField("Owner", "", 20, nil, nil)
	form.AddPasswordField("CVV", "", 20, '*', nil)
	form.AddPasswordField("Secret key", "", 20, '*', nil)
	form.AddCheckbox("Save local", false, nil)
	form.AddCheckbox("Save remote", false, nil)
	form.AddButton("Save", func() {

	})

	flex.AddItem(infoBox, 0, 1, false)
	flex.AddItem(form, 0, 4, false)
	flex.SetDirection(0)

	return flex
}
