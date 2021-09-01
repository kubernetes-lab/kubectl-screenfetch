package ascii

import (
	"errors"
	"os"

	"golang.org/x/term"
)

const (
	charWidth = 0.5
)

type Terminal struct {
}

func NewTerminal() *Terminal {
	return &Terminal{}
}

func (t *Terminal) CharWidth() float64 {
	return charWidth
}

func (t *Terminal) ScreenSize() (width, height int, err error) {
	if !term.IsTerminal(int(os.Stdout.Fd())) {
		return 0, 0, errors.New("Terminal is not detected")
	}

	w, h, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 0, 0, err
	}

	return w, h, nil
}
