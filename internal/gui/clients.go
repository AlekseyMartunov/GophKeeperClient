package gui

import (
	"fmt"
	"github.com/rivo/tview"
)

func newClientsPage(tokenClient tokenClientHTTP) *tview.Flex {
	flex := tview.NewFlex()
	errBlock := tview.NewTextView()
	getClientButton := tview.NewButton("Get your clients")
	list := tview.NewList()
	blockTokenButton := tview.NewButton("Block")
	clients := make([]string, 0)

	blockTokenButton.SetBorder(true)

	list.SetChangedFunc(func(index int, name string, secondName string, shortcut rune) {
		blockTokenButton.SetLabel(fmt.Sprintf("Block %s", clients[index]))
		blockTokenButton.SetSelectedFunc(func() {
			err := tokenClient.BlockToken(clients[index])
			if err != nil {
				errBlock.Clear()
				errBlock.SetText(err.Error())
			} else {
				errBlock.Clear()
				errBlock.SetText(fmt.Sprintf("blocked successfully %s", clients[index]))
			}
		})
	}).SetWrapAround(true)

	getClientButton.SetBorder(true)
	getClientButton.SetSelectedFunc(func() {
		c, err := tokenClient.GetAll()
		if err != nil {
			errBlock.Clear()
			errBlock.SetText(err.Error())
		} else {
			clients = append(clients, c...)
			for _, clientName := range c {
				list.AddItem(clientName, "", rune('*'), nil)
			}
		}

	})

	flex.SetDirection(tview.FlexRow)
	flex.AddItem(
		tview.NewFlex().
			SetDirection(tview.FlexColumn).
			AddItem(getClientButton, 0, 1, false).
			AddItem(blockTokenButton, 0, 1, false).
			AddItem(errBlock, 0, 3, false),
		0, 1, false)
	flex.AddItem(list, 0, 9, false)

	return flex
}
