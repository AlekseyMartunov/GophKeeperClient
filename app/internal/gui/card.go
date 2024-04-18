package gui

import (
	card "GophKeeperClient/internal/entity/card"
	"fmt"
	"github.com/rivo/tview"
	"time"
)

func newCardPageAdd(cardService cardService) *tview.Flex {
	flex := tview.NewFlex()
	form := tview.NewForm()
	infoBox := tview.NewTextView()
	errBox := tview.NewTextView()

	infoBox.SetText("Add credit card ...")
	infoBox.SetBorderPadding(2, 2, 2, 2)

	form.AddInputField("Name", "", 20, nil, nil)
	form.AddInputField("Number", "", 20, nil, nil)
	form.AddInputField("Date", "", 20, nil, nil)
	form.AddInputField("Owner", "", 20, nil, nil)
	form.AddPasswordField("CVV", "", 20, '*', nil)
	form.AddPasswordField("Secret key", "", 20, '*', nil)
	form.AddPasswordField("Secret key repeat", "", 20, '*', nil)
	form.AddCheckbox("Save local", false, nil)
	form.AddCheckbox("Save remote", false, nil)
	form.AddButton("Save", func() {
		name := form.GetFormItem(0).(*tview.InputField).GetText()
		number := form.GetFormItem(1).(*tview.InputField).GetText()
		date := form.GetFormItem(2).(*tview.InputField).GetText()
		owner := form.GetFormItem(3).(*tview.InputField).GetText()
		cvv := form.GetFormItem(4).(*tview.InputField).GetText()
		secretKey := form.GetFormItem(5).(*tview.InputField).GetText()
		secretKeyRepeat := form.GetFormItem(6).(*tview.InputField).GetText()
		saveLocal := form.GetFormItem(7).(*tview.Checkbox).IsChecked()
		saveRemote := form.GetFormItem(8).(*tview.Checkbox).IsChecked()

		if secretKeyRepeat != secretKey {
			errBox.SetText("secret keys don't match")
		} else {
			errBox.Clear()

			c := card.Card{
				Name:        name,
				Number:      number,
				CVV:         cvv,
				Owner:       owner,
				Date:        date,
				CreatedTime: time.Now(),
			}

			if saveLocal {
				err := cardService.SaveLocal(c, secretKey)
				if err != nil {
					errBox.SetText(err.Error())
				} else {
					errBox.SetText("save local successfully")
				}
			}

			if saveRemote {
				err := cardService.SaveRemote(c, secretKey)
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
	flex.AddItem(form, 0, 7, false)
	flex.AddItem(errBox, 0, 1, false)
	flex.SetDirection(tview.FlexRow)

	return flex
}
