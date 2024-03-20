package gui

import "github.com/rivo/tview"

func newPasswordPage() *tview.Flex {
	flex := tview.NewFlex()
	form := tview.NewForm()
	infoBox := tview.NewTextArea()

	infoBox.SetText("Add password ...", false)
	infoBox.SetBorderPadding(2, 2, 2, 2)

	form.AddInputField("Name", "", 20, nil, nil)
	form.AddInputField("Login", "", 20, nil, nil)
	form.AddPasswordField("Password", "", 20, '*', nil)
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
