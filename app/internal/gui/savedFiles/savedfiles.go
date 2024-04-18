package savedfiles

import (
	"GophKeeperClient/internal/entity/file"
	"fmt"
	"github.com/rivo/tview"
	"os"
	"path"
	"time"
)

type fileService interface {
	SaveLocal(f *file.File) error
	SaveRemote(f *file.File) error
	GetFromLocal(fileName string) (*file.File, error)
	GetFromRemote(fileName string) (*file.File, error)
	GetAllFromLocal() ([]*file.File, error)
	GetAllFromRemote() ([]*file.File, error)
	DeleteFromLocal(fileName string) error
	DeleteFromRemote(fileName string) error
}

var errorBox = tview.NewTextView()
var getDetailInfoButton = tview.NewButton("Get ...")
var detailInfoBox = tview.NewTextView()

var localFilesList = tview.NewList()
var localFiles = make([]*file.File, 0)
var removeLocalFilesButton = tview.NewButton("Remove ...")

var remoteFilesList = tview.NewList()
var remoteFiles = make([]*file.File, 0)
var removeRemoteFileButton = tview.NewButton("Remove ...")

func NewSavedPairsPage(service fileService) *tview.Flex {
	mainFlex := tview.NewFlex()
	localFlex := tview.NewFlex()
	remoteFlex := tview.NewFlex()
	infoFlex := tview.NewFlex()

	mainFlex.AddItem(localFlex, 0, 1, false)
	mainFlex.AddItem(remoteFlex, 0, 1, false)
	mainFlex.AddItem(infoFlex, 0, 1, false)
	mainFlex.SetDirection(tview.FlexColumn)

	localFlex.AddItem(tview.NewButton("Get local files").SetSelectedFunc(func() {
		loadLocalCards(service)
	}), 0, 1, false)
	localFlex.AddItem(removeLocalFilesButton, 0, 1, false)
	localFlex.AddItem(localFilesList, 0, 4, false)
	localFlex.SetDirection(tview.FlexRow)

	localFilesList.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		selectedCard := localFiles[index].Name
		updateRemoveLocalButton(service, selectedCard)
		updateGetDetailInfoButtonLocal(service, selectedCard)
	})

	remoteFlex.AddItem(tview.NewButton("Get remote files").SetSelectedFunc(func() {
		loadRemoteCards(service)
	}), 0, 1, false)
	remoteFlex.AddItem(removeRemoteFileButton, 0, 1, false)
	remoteFlex.AddItem(remoteFilesList, 0, 4, false)
	remoteFlex.SetDirection(tview.FlexRow)

	remoteFilesList.SetSelectedFunc(func(i int, s string, s2 string, r rune) {
		selectedCard := remoteFiles[i].Name
		updateRemoveRemoteButton(service, selectedCard)
		updateGetDetailInfoButtonRemote(service, selectedCard)
	})

	infoFlex.AddItem(getDetailInfoButton, 0, 1, false)
	infoFlex.AddItem(detailInfoBox, 0, 2, false)
	infoFlex.AddItem(errorBox, 0, 2, false)
	infoFlex.SetDirection(tview.FlexRow)

	return mainFlex
}

func loadLocalCards(service fileService) {
	localFiles = localFiles[:0]
	localFilesList.Clear()

	files, err := service.GetAllFromLocal()
	if err != nil {
		errorBox.Clear()
		errorBox.SetText(err.Error())
	} else {
		localFiles = append(localFiles, files...)
		for _, c := range files {
			localFilesList.AddItem(c.Name, "", '*', nil)
		}
	}
}

func loadRemoteCards(service fileService) {
	remoteFiles = remoteFiles[:0]
	remoteFilesList.Clear()

	files, err := service.GetAllFromRemote()
	if err != nil {
		errorBox.Clear()
		errorBox.SetText(err.Error())
	} else {
		remoteFiles = append(remoteFiles, files...)
		for _, c := range files {
			remoteFilesList.AddItem(c.Name, "", '*', nil)
		}
	}
}

func updateRemoveRemoteButton(service fileService, name string) {
	removeRemoteFileButton.SetLabel(fmt.Sprintf("Remove %s", name))
	removeRemoteFileButton.SetSelectedFunc(func() {
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

func updateRemoveLocalButton(service fileService, name string) {
	removeLocalFilesButton.SetLabel(fmt.Sprintf("Remove %s", name))
	removeLocalFilesButton.SetSelectedFunc(func() {
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

func updateGetDetailInfoButtonLocal(service fileService, name string) {
	getDetailInfoButton.SetLabel(fmt.Sprintf("Get detail info %s", name))
	getDetailInfoButton.SetSelectedFunc(func() {
		file, err := service.GetFromLocal(name)

		if err != nil {
			errorBox.Clear()
			errorBox.SetText(err.Error())
		} else {

			err = createFile(file)
			if err != nil {
				errorBox.SetText(err.Error())
				return
			}
			detailInfoBox.SetText("file created successfully ")
		}
	})
}

func updateGetDetailInfoButtonRemote(service fileService, name string) {
	getDetailInfoButton.SetLabel(fmt.Sprintf("Get detail info %s", name))
	getDetailInfoButton.SetSelectedFunc(func() {
		file, err := service.GetFromRemote(name)

		if err != nil {
			errorBox.Clear()
			errorBox.SetText(err.Error())
		} else {

			err = createFile(file)
			if err != nil {
				errorBox.SetText(err.Error())
				return
			}
			detailInfoBox.SetText("file created successfully ")
		}
	})
}

func createFile(f *file.File) error {
	t := time.Now()
	s := fmt.Sprintf("%s-%s", t.Format(time.DateOnly), t.Format(time.TimeOnly))

	err := os.Mkdir(s, 0777)
	if err != nil {
		return err
	}

	createdFile, err := os.Create(path.Join(s, f.Name))
	if err != nil {
		return err
	}

	defer createdFile.Close()

	_, err = createdFile.Write(f.Data)
	if err != nil {
		return err
	}

	return nil
}
