package gui

import (
	"github.com/rivo/tview"
)

func Run() error {
	app := tview.NewApplication()

	pages := tview.NewPages()
	pages.AddPage("login", newLoginPage(), true, false)
	pages.AddPage("registration", newRegistrationPage(), true, false)
	pages.AddPage("clients", newClientsPage(), true, false)
	pages.AddPage("card", newCardPage(), true, false)
	pages.AddPage("password", newPasswordPage(), true, false)

	list := tview.NewList()

	list.AddItem("LOGIN", "", '1', func() {
		pages.SwitchToPage("login")
	})

	list.AddItem("REGISTRATION", "", '2', func() {
		pages.SwitchToPage("registration")
	})

	list.AddItem("CLIENTS", "", '3', func() {
		pages.SwitchToPage("clients")
	})

	list.AddItem("ADD CREDIT CARD", "", '4', func() {
		pages.SwitchToPage("card")
	})

	list.AddItem("ADD PASSWORD", "", '5', func() {
		pages.SwitchToPage("password")
	})

	list.AddItem("ADD FILE", "", '6', func() {
		pages.SwitchToPage("file")
	})

	list.AddItem("LOCAl DATA", "", '6', func() {
		pages.SwitchToPage("local")
	})

	list.AddItem("REMOTE DATA", "", '6', func() {
		pages.SwitchToPage("remote")
	})

	list.AddItem("EXIT", "", '7', func() {
		app.Stop()
	})

	list.SetBorder(true)
	list.SetTitle("MENU")

	flex := tview.NewFlex().
		AddItem(list, 0, 1, true).
		AddItem(pages, 0, 5, false)
	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		return err
	}

	return nil
}
