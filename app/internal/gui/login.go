package gui

import (
	"github.com/rivo/tview"
)

func newLoginPage(tokenClient tokenClientHTTP) *tview.Flex {
	infoBox := tview.NewTextArea()

	form := tview.NewForm()
	form.AddInputField("your login", "", 20, nil, nil)
	form.AddPasswordField("your password", "", 20, '*', nil)
	form.AddInputField("name of this client", "", 20, nil, nil)

	form.AddButton("LOGIN", func() {
		login := form.GetFormItem(0).(*tview.InputField).GetText()
		password := form.GetFormItem(1).(*tview.InputField).GetText()
		clientName := form.GetFormItem(2).(*tview.InputField).GetText()

		err := tokenClient.UpdateToken(clientName, login, password)
		if err != nil {
			infoBox.SetText(err.Error(), false)
		} else {
			infoBox.SetText("OK", false)
		}
	})

	flex := tview.NewFlex()
	flex.AddItem(form, 0, 1, false)
	flex.AddItem(infoBox, 0, 1, false)
	flex.SetDirection(0)

	return flex
}
