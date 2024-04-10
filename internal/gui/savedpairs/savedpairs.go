package savedpairs

import (
	"GophKeeperClient/internal/entity/pair"
	"fmt"
	"github.com/rivo/tview"
)

type passwordService interface {
	SaveLocal(p pair.Pair, key string) error
	SaveRemote(p pair.Pair, key string) error
	GetFromLocal(name, key string) (pair.Pair, error)
	GetFromRemote(name, key string) (pair.Pair, error)
	GetAllFromLocal() ([]pair.Pair, error)
	GetAllFromRemote() ([]pair.Pair, error)
	DeleteFromLocal(name string) error
	DeleteFromRemote(name string) error
}

var errorBox = tview.NewTextView()
var getDetailInfoButton = tview.NewButton("Get ...")
var secretKeyForm = tview.NewForm()
var detailInfoBox = tview.NewTextView()

var localPairsList = tview.NewList()
var localPairs = make([]pair.Pair, 0)
var removeLocalPairButton = tview.NewButton("Remove ...")

var remotePairsList = tview.NewList()
var remotePairs = make([]pair.Pair, 0)
var removeRemotePairButton = tview.NewButton("Remove ...")

func NewSavedPairsPage(service passwordService) *tview.Flex {
	mainFlex := tview.NewFlex()
	localFlex := tview.NewFlex()
	remoteFlex := tview.NewFlex()
	infoFlex := tview.NewFlex()

	mainFlex.AddItem(localFlex, 0, 1, false)
	mainFlex.AddItem(remoteFlex, 0, 1, false)
	mainFlex.AddItem(infoFlex, 0, 1, false)
	mainFlex.SetDirection(tview.FlexColumn)

	localFlex.AddItem(tview.NewButton("Get local pairs").SetSelectedFunc(func() {
		loadLocalCards(service)
	}), 0, 1, false)
	localFlex.AddItem(removeLocalPairButton, 0, 1, false)
	localFlex.AddItem(localPairsList, 0, 4, false)
	localFlex.SetDirection(tview.FlexRow)

	localPairsList.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		selectedCard := localPairs[index].Name
		updateRemoveLocalButton(service, selectedCard)
		updateGetDetailInfoButtonLocal(service, selectedCard)
	})

	remoteFlex.AddItem(tview.NewButton("Get remote pairs").SetSelectedFunc(func() {
		loadRemoteCards(service)
	}), 0, 1, false)
	remoteFlex.AddItem(removeRemotePairButton, 0, 1, false)
	remoteFlex.AddItem(remotePairsList, 0, 4, false)
	remoteFlex.SetDirection(tview.FlexRow)

	remotePairsList.SetSelectedFunc(func(i int, s string, s2 string, r rune) {
		selectedCard := remotePairs[i].Name
		updateRemoveRemoteButton(service, selectedCard)
		updateGetDetailInfoButtonRemote(service, selectedCard)
	})

	secretKeyForm.AddPasswordField("secret key", "", 20, '*', nil)

	infoFlex.AddItem(getDetailInfoButton, 0, 1, false)
	infoFlex.AddItem(secretKeyForm, 0, 1, false)
	infoFlex.AddItem(detailInfoBox, 0, 2, false)
	infoFlex.AddItem(errorBox, 0, 1, false)
	infoFlex.SetDirection(tview.FlexRow)

	return mainFlex
}

func loadLocalCards(service passwordService) {
	localPairs = localPairs[:0]
	localPairsList.Clear()

	pairs, err := service.GetAllFromLocal()
	if err != nil {
		errorBox.Clear()
		errorBox.SetText(err.Error())
	} else {
		localPairs = append(localPairs, pairs...)
		for _, c := range pairs {
			localPairsList.AddItem(c.Name, "", '*', nil)
		}
	}
}

func loadRemoteCards(service passwordService) {
	remotePairs = remotePairs[:0]
	remotePairsList.Clear()

	pairs, err := service.GetAllFromRemote()
	if err != nil {
		errorBox.Clear()
		errorBox.SetText(err.Error())
	} else {
		remotePairs = append(remotePairs, pairs...)
		for _, c := range pairs {
			remotePairsList.AddItem(c.Name, "", '*', nil)
		}
	}
}

func updateRemoveRemoteButton(service passwordService, name string) {
	removeRemotePairButton.SetLabel(fmt.Sprintf("Remove %s", name))
	removeRemotePairButton.SetSelectedFunc(func() {
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

func updateRemoveLocalButton(service passwordService, name string) {
	removeLocalPairButton.SetLabel(fmt.Sprintf("Remove %s", name))
	removeLocalPairButton.SetSelectedFunc(func() {
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

func updateGetDetailInfoButtonLocal(service passwordService, name string) {
	getDetailInfoButton.SetLabel(fmt.Sprintf("Get detail info %s", name))
	getDetailInfoButton.SetSelectedFunc(func() {
		key := secretKeyForm.GetFormItem(0).(*tview.InputField).GetText()
		pair, err := service.GetFromLocal(name, key)

		if err != nil {
			errorBox.Clear()
			errorBox.SetText(err.Error())
		} else {
			s := fmt.Sprintf("Name: %s\nLogin: %s\nPassword: %s\nCreated time: %s",
				pair.Name, pair.Login, pair.Password, pair.CreatedTime.String())
			detailInfoBox.SetText(s)
		}
	})
}

func updateGetDetailInfoButtonRemote(service passwordService, name string) {
	getDetailInfoButton.SetLabel(fmt.Sprintf("Get detail info %s", name))
	getDetailInfoButton.SetSelectedFunc(func() {
		key := secretKeyForm.GetFormItem(0).(*tview.InputField).GetText()
		pair, err := service.GetFromRemote(name, key)

		if err != nil {
			errorBox.Clear()
			errorBox.SetText(err.Error())
		} else {
			s := fmt.Sprintf("Name: %s\nLogin: %s\nPassword: %s\nCreated time: %s",
				pair.Name, pair.Login, pair.Password, pair.CreatedTime.String())
			detailInfoBox.SetText(s)
		}
	})
}
