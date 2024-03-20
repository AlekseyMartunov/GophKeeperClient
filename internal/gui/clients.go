package gui

import (
	"crypto/rand"
	"fmt"
	"github.com/rivo/tview"
	"strings"
)

// iterator support struct for iteration input data
type iterator struct {
	tokens   []string
	current  int
	maxIndex int
}

func newIterator(s []string) *iterator {
	return &iterator{
		tokens:   s,
		current:  0,
		maxIndex: len(s) - 1,
	}
}

func (i *iterator) IsAvailable() bool {
	return i.current <= i.maxIndex
}

func (i *iterator) Next() string {
	res := i.tokens[i.current]
	if i.current <= i.maxIndex {
		i.current += 1
	}
	return res
}

func newClientsPage() *tview.Flex {
	infoText := `here are all your clients you can block selected`
	infoBox := tview.NewTextArea()
	infoBox.SetText(infoText, false)
	infoBox.SetBorderPadding(1, 2, 3, 4)

	tokenForBlocking := make(map[string]bool)
	iter := newIterator(getToken())

	form := tview.NewForm()
	flex := tview.NewFlex()

	for iter.IsAvailable() {
		token := iter.Next()
		form.AddCheckbox(token, false, func(checked bool) {
			tokenForBlocking[token] = checked
		})
	}

	form.AddButton("BLOCK", func() {
		str := make([]string, 5)
		for k, v := range tokenForBlocking {
			str = append(str, fmt.Sprintf("%s - %t", k, v))
		}
		infoBox.SetText(strings.Join(str, "\n"), false)
	})

	flex.AddItem(infoBox, 0, 1, false)
	flex.AddItem(form, 0, 1, false)
	flex.SetDirection(0)

	return flex

}

func getToken() []string {
	b := make([]byte, 10)
	rand.Read(b)

	return []string{string(b), "b", "c", "d", "f"}
}
