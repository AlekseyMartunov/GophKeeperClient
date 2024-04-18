package gui

import (
	"GophKeeperClient/internal/entity/pair"
	"fmt"
	"github.com/rivo/tview"
	"time"
)

func newPasswordPage(password passwordService) *tview.Flex {
	flex := tview.NewFlex()
	form := tview.NewForm()
	infoBox := tview.NewTextView()
	errBox := tview.NewTextView()

	infoBox.SetText("Add password ...")
	infoBox.SetBorderPadding(2, 2, 2, 2)

	form.AddInputField("Name", "", 20, nil, nil)
	form.AddInputField("Login", "", 20, nil, nil)
	form.AddPasswordField("Password", "", 20, '*', nil)
	form.AddPasswordField("Secret key", "", 20, '*', nil)
	form.AddPasswordField("Secret key repeat", "", 20, '*', nil)
	form.AddCheckbox("Save local", false, nil)
	form.AddCheckbox("Save remote", false, nil)
	form.AddButton("Save", func() {
		name := form.GetFormItem(0).(*tview.InputField).GetText()
		login := form.GetFormItem(1).(*tview.InputField).GetText()
		pass := form.GetFormItem(2).(*tview.InputField).GetText()
		secretKey := form.GetFormItem(3).(*tview.InputField).GetText()
		keyRepeat := form.GetFormItem(4).(*tview.InputField).GetText()
		saveLocal := form.GetFormItem(5).(*tview.Checkbox).IsChecked()
		saveRemote := form.GetFormItem(6).(*tview.Checkbox).IsChecked()

		if secretKey != keyRepeat {
			errBox.SetText("secret keys don't match")
		} else {
			errBox.Clear()

			p := pair.Pair{
				Name:        name,
				Login:       login,
				Password:    pass,
				CreatedTime: time.Now(),
			}

			if saveLocal {
				err := password.SaveLocal(p, secretKey)
				if err != nil {
					errBox.SetText(err.Error())
				} else {
					errBox.SetText("save local successfully")
				}
			}

			if saveRemote {
				err := password.SaveRemote(p, secretKey)
				if err != nil {
					text := errBox.GetText(true)
					errBox.SetText(fmt.Sprintf("%s \n %s", text, err.Error()))
				} else {
					text := errBox.GetText(true)
					errBox.SetText(fmt.Sprintf("%s \n %s", text, "save remote successfully"))
				}
			}
		}
	})

	flex.AddItem(infoBox, 0, 1, false)
	flex.AddItem(form, 0, 6, false)
	flex.AddItem(errBox, 0, 3, false)
	flex.SetDirection(tview.FlexRow)

	return flex
}
