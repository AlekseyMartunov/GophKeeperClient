package gui

import (
	"github.com/rivo/tview"
)

func newRegistrationPage(userClient userClientHTTP) *tview.Flex {
	infoBox := tview.NewTextArea()

	form := tview.NewForm()
	form.AddInputField("enter your login", "", 20, nil, nil)
	form.AddPasswordField("enter your password", "", 20, '*', nil)
	form.AddPasswordField("enter your password agan", "", 20, '*', nil)
	form.AddButton("REGISTRATION", func() {
		Login := form.GetFormItem(0).(*tview.InputField).GetText()
		password := form.GetFormItem(1).(*tview.InputField).GetText()
		passwordRepeat := form.GetFormItem(2).(*tview.InputField).GetText()

		if password != passwordRepeat {
			infoBox.SetText("Пароли не совпадают", false)
		} else {
			err := userClient.RegisterUser(Login, password)
			if err != nil {
				infoBox.SetText(err.Error(), false)
			} else {
				infoBox.SetText("OK", false)
			}
		}
	})

	flex := tview.NewFlex()
	flex.AddItem(form, 0, 1, false)
	flex.AddItem(infoBox, 0, 1, false)
	flex.SetDirection(0)

	return flex
}
