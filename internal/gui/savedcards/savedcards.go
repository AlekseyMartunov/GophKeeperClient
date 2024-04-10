package savedcards

import (
	"GophKeeperClient/internal/entity/card"
	"fmt"
	"github.com/rivo/tview"
)

type cardService interface {
	SaveLocal(c card.Card, key string) error
	SaveRemote(c card.Card, key string) error
	GetFromLocal(name, key string) (card.Card, error)
	GetFromRemote(name, key string) (card.Card, error)
	GetAllFromLocal() ([]card.Card, error)
	GetAllFromRemote() ([]card.Card, error)
	DeleteFromLocal(name string) error
	DeleteFromRemote(name string) error
}

var errorBox = tview.NewTextView()
var getDetailInfoButton = tview.NewButton("Get ...")
var secretKeyForm = tview.NewForm()
var detailInfoBox = tview.NewTextView()

var localCardsList = tview.NewList()
var localCards = make([]card.Card, 0)
var removeLocalCardButton = tview.NewButton("Remove ...")

var remoteCardsList = tview.NewList()
var remoteCards = make([]card.Card, 0)
var removeRemoteCardButton = tview.NewButton("Remove ...")

func NewSavedCardsPage(cardService cardService) *tview.Flex {
	mainFlex := tview.NewFlex()
	localFlex := tview.NewFlex()
	remoteFlex := tview.NewFlex()
	infoFlex := tview.NewFlex()

	mainFlex.AddItem(localFlex, 0, 1, false)
	mainFlex.AddItem(remoteFlex, 0, 1, false)
	mainFlex.AddItem(infoFlex, 0, 1, false)
	mainFlex.SetDirection(tview.FlexColumn)

	localFlex.AddItem(tview.NewButton("Get local cards").SetSelectedFunc(func() {
		loadLocalCards(cardService)
	}), 0, 1, false)
	localFlex.AddItem(removeLocalCardButton, 0, 1, false)
	localFlex.AddItem(localCardsList, 0, 4, false)
	localFlex.SetDirection(tview.FlexRow)

	localCardsList.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		selectedCard := localCards[index].Name
		updateRemoveLocalButton(cardService, selectedCard)
		updateGetDetailInfoButtonLocal(cardService, selectedCard)
	})

	remoteFlex.AddItem(tview.NewButton("Get remote cards").SetSelectedFunc(func() {
		loadRemoteCards(cardService)
	}), 0, 1, false)
	remoteFlex.AddItem(removeRemoteCardButton, 0, 1, false)
	remoteFlex.AddItem(remoteCardsList, 0, 4, false)
	remoteFlex.SetDirection(tview.FlexRow)

	remoteCardsList.SetSelectedFunc(func(i int, s string, s2 string, r rune) {
		selectedCard := remoteCards[i].Name
		updateRemoveRemoteButton(cardService, selectedCard)
		updateGetDetailInfoButtonRemote(cardService, selectedCard)
	})

	secretKeyForm.AddPasswordField("secret key", "", 20, '*', nil)

	infoFlex.AddItem(getDetailInfoButton, 0, 1, false)
	infoFlex.AddItem(secretKeyForm, 0, 1, false)
	infoFlex.AddItem(detailInfoBox, 0, 2, false)
	infoFlex.AddItem(errorBox, 0, 1, false)
	infoFlex.SetDirection(tview.FlexRow)

	return mainFlex
}

func loadLocalCards(cardService cardService) {
	localCards = localCards[:0]
	localCardsList.Clear()

	cards, err := cardService.GetAllFromLocal()
	if err != nil {
		errorBox.Clear()
		errorBox.SetText(err.Error())
	} else {
		localCards = append(localCards, cards...)
		for _, c := range cards {
			localCardsList.AddItem(c.Name, "", '*', nil)
		}
	}
}

func loadRemoteCards(cardService cardService) {
	remoteCards = remoteCards[:0]
	remoteCardsList.Clear()

	cards, err := cardService.GetAllFromRemote()
	if err != nil {
		errorBox.Clear()
		errorBox.SetText(err.Error())
	} else {
		remoteCards = append(remoteCards, cards...)
		for _, c := range cards {
			remoteCardsList.AddItem(c.Name, "", '*', nil)
		}
	}
}

func updateRemoveRemoteButton(service cardService, name string) {
	removeRemoteCardButton.SetLabel(fmt.Sprintf("Remove %s", name))
	removeRemoteCardButton.SetSelectedFunc(func() {
		err := service.DeleteFromRemote(name)
		if err != nil {
			errorBox.Clear()
			errorBox.SetText(err.Error())
		} else {
			errorBox.Clear()
			errorBox.SetText(fmt.Sprintf("%s deleted successfully", name))
			loadRemoteCards(service)
		}
	})
}

func updateRemoveLocalButton(service cardService, name string) {
	removeLocalCardButton.SetLabel(fmt.Sprintf("Remove %s", name))
	removeLocalCardButton.SetSelectedFunc(func() {
		err := service.DeleteFromLocal(name)
		if err != nil {
			errorBox.Clear()
			errorBox.SetText(err.Error())
		} else {
			errorBox.Clear()
			errorBox.SetText(fmt.Sprintf("%s deleted successfully", name))
			loadLocalCards(service)
		}
	})
}

func updateGetDetailInfoButtonLocal(service cardService, name string) {
	getDetailInfoButton.SetLabel(fmt.Sprintf("Get detail info %s", name))
	getDetailInfoButton.SetSelectedFunc(func() {
		key := secretKeyForm.GetFormItem(0).(*tview.InputField).GetText()
		card, err := service.GetFromLocal(name, key)

		if err != nil {
			errorBox.Clear()
			errorBox.SetText(err.Error())
		} else {
			s := fmt.Sprintf("Name: %s\nNumber: %s\nOwner: %s\nDate: %s\nCVV: %s\nCreated time: %s",
				card.Name, card.Number, card.Owner, card.Date, card.CVV, card.CreatedTime.String())
			detailInfoBox.SetText(s)
		}
	})
}

func updateGetDetailInfoButtonRemote(service cardService, name string) {
	getDetailInfoButton.SetLabel(fmt.Sprintf("Get detail info %s", name))
	getDetailInfoButton.SetSelectedFunc(func() {
		key := secretKeyForm.GetFormItem(0).(*tview.InputField).GetText()
		card, err := service.GetFromRemote(name, key)

		if err != nil {
			errorBox.Clear()
			errorBox.SetText(err.Error())
		} else {
			s := fmt.Sprintf("Name: %s\nNumber: %s\nOwner: %s\nDate: %s\nCVV: %s\nCreated time: %s",
				card.Name, card.Number, card.Owner, card.Date, card.CVV, card.CreatedTime.String())
			detailInfoBox.SetText(s)
		}
	})
}
