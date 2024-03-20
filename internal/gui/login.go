package gui

import (
	"fmt"
	"github.com/rivo/tview"
)

func newLoginPage() *tview.Flex {

	infoBox := tview.NewTextArea()
	infoBox.SetText("some text here", false)

	form := tview.NewForm()
	form.AddInputField("your login", "", 20, nil, nil)
	form.AddPasswordField("your password", "", 20, '*', nil)
	form.AddButton("LOGIN", func() {
		textLogin := form.GetFormItem(0).(*tview.InputField).GetText()
		textPassword := form.GetFormItem(1).(*tview.InputField).GetText()
		infoBox.SetText(fmt.Sprintf("%s \n %s", textLogin, textPassword), false)
	})

	flex := tview.NewFlex()
	flex.AddItem(form, 0, 1, false)
	flex.AddItem(infoBox, 0, 1, false)
	flex.SetDirection(0)

	return flex
}
