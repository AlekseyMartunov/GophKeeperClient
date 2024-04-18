package gui

import (
	"GophKeeperClient/internal/entity/file"
	"fmt"
	"github.com/rivo/tview"
	"os"
	"time"
)

func NewAddFilePage(service fileService) *tview.Flex {
	flex := tview.NewFlex()
	form := tview.NewForm()
	errBox := tview.NewTextView()
	infoBox := tview.NewTextView()
	fileList := tview.NewList()

	form.AddInputField("set file path", "", 40, nil, nil)
	form.AddCheckbox("Save local", false, nil)
	form.AddCheckbox("Save remote", false, nil)

	button := tview.NewButton("Open").SetSelectedFunc(func() {
		fileList.Clear()
		p := form.GetFormItem(0).(*tview.InputField).GetText()

		f, err := os.Open(p)
		defer f.Close()

		if err != nil {
			errBox.SetText(err.Error())
			return
		}

		info, err := f.Stat()
		if err != nil {
			errBox.SetText(err.Error())
			return
		}

		if info.IsDir() {
			dirs, err := os.ReadDir(p)
			if err != nil {
				errBox.SetText(err.Error())
				return
			}
			for _, d := range dirs {
				fileList.AddItem(d.Name(), "", '*', nil)
			}
			return
		}
		b, err := os.ReadFile(p)
		if err != nil {
			errBox.SetText(err.Error())
		}
		errBox.SetText(fmt.Sprintf("read %d bytes, path: %s", len(b), p))

		saveLocal := form.GetFormItem(1).(*tview.Checkbox).IsChecked()
		saveRemote := form.GetFormItem(2).(*tview.Checkbox).IsChecked()

		fileEntity := file.File{
			Name:        info.Name(),
			Data:        b,
			CreatedTime: time.Now(),
		}

		if saveLocal {
			err = service.SaveLocal(&fileEntity)
			if err != nil {
				errBox.SetText(err.Error())
			}
		}

		if saveRemote {
			err = service.SaveRemote(&fileEntity)
			if err != nil {
				errBox.SetText(err.Error())
			}
		}

		infoBox.SetText(fmt.Sprintf("file successfully readed: %s", info.Name()))

	})

	flex.AddItem(infoBox, 0, 1, false)
	flex.AddItem(fileList, 0, 5, false)
	flex.AddItem(errBox, 0, 1, false)
	flex.AddItem(form, 0, 4, false)
	flex.AddItem(button, 0, 1, false)
	flex.SetDirection(tview.FlexRow)

	return flex
}
