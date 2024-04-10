package gui

import (
	savedfiles "GophKeeperClient/internal/gui/savedFiles"
	"GophKeeperClient/internal/gui/savedcards"
	"GophKeeperClient/internal/gui/savedpairs"
	"github.com/rivo/tview"
)

func Run(cfg config,
	userClient userClientHTTP,
	tokenClient tokenClientHTTP,
	card cardService,
	password passwordService,
	file fileService,
) error {
	app := tview.NewApplication()

	pages := tview.NewPages()
	pages.AddPage("registration", newRegistrationPage(userClient), true, false)
	pages.AddPage("login", newLoginPage(tokenClient), true, false)
	pages.AddPage("clients", newClientsPage(tokenClient), true, false)
	pages.AddPage("card", newCardPageAdd(card), true, false)
	pages.AddPage("password", newPasswordPage(password), true, false)
	pages.AddPage("addFile", NewAddFilePage(file), true, false)
	pages.AddPage("savedCards", savedcards.NewSavedCardsPage(card), true, false)
	pages.AddPage("savedPairs", savedpairs.NewSavedPairsPage(password), true, false)
	pages.AddPage("savedFiles", savedfiles.NewSavedPairsPage(file), true, false)

	list := tview.NewList()

	list.AddItem("REGISTRATION", "", '1', func() {
		pages.SwitchToPage("registration")
	})

	list.AddItem("LOGIN", "", '2', func() {
		pages.SwitchToPage("login")
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
		pages.SwitchToPage("addFile")
	})

	list.AddItem("SAVED CREDIT CARDS", "", '7', func() {
		pages.SwitchToPage("savedCards")
	})

	list.AddItem("SAVED PASSWORDS", "", '8', func() {
		pages.SwitchToPage("savedPairs")
	})

	list.AddItem("SAVED FILES", "", '9', func() {
		pages.SwitchToPage("savedFiles")
	})

	list.AddItem("EXIT", "", '*', func() {
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
