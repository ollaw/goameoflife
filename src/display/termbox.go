package display

import (
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
	"gitlab.com/ollaww/goameoflife/src/matrix"
)

type TermboxDisplayer struct {
}

const (
	FOOTER_FG      = termbox.ColorWhite
	FOOTER_BG      = termbox.ColorDefault
	FOOTER_TEXT    = "Round: %d * Births: %d * Deaths: %d * Press Esc to quit"
	CELL_TEXT      = ' '
	SEPARATOR_TEXT = '|'
)

func (i *TermboxDisplayer) Init(m *matrix.Matrix) error {
	if err := termbox.Init(); err != nil {
		return err
	}

	// Bind Esc to quit
	go func() {
		for {
			ev := termbox.PollEvent()
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				termbox.Close()
				os.Exit(0)
			}
		}
	}()
	displayFooter(m)
	return nil
}

func (i *TermboxDisplayer) Display(m *matrix.Matrix) error {
	i.Clear(m)

	// Center matrix
	baseRowOffset, baseColOffset := baseOffsets(m)
	// Additional line separator
	additionalrowOffset := 0

	for k := range *m.Matrix {
		// Used for SEPARATOR_TEXT sparator
		additionalColOffset := 0
		for j := range (*m.Matrix)[k] {
			termbox.SetCell(baseColOffset+j+additionalColOffset, baseRowOffset+k+additionalrowOffset, SEPARATOR_TEXT, termbox.ColorDefault, termbox.ColorDefault)
			additionalColOffset++
			cellColor := termbox.ColorWhite
			if (*m.Matrix)[k][j] {
				cellColor = termbox.ColorGreen
			}
			termbox.SetCell(baseColOffset+j+additionalColOffset, baseRowOffset+k+additionalrowOffset, CELL_TEXT, termbox.ColorDefault, cellColor)
		}
		additionalrowOffset++
	}
	displayFooter(m)
	termbox.Flush()
	return nil
}

func (i *TermboxDisplayer) Update(m *matrix.Matrix) error {
	return i.Display(m)
}
func (i *TermboxDisplayer) Clear(m *matrix.Matrix) error {
	return termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

// Compute the base coordinates of matrix into terminal
func baseOffsets(m *matrix.Matrix) (int, int) {
	sizeW, sizeH := termbox.Size()
	row := (sizeH / 2) - len(*m.Matrix)
	col := (sizeW / 2) - len((*m.Matrix)[0])
	return row, col
}

func displayFooter(m *matrix.Matrix) {
	sizeW, sizeH := termbox.Size()
	tbprint((sizeW/2)-len(FOOTER_TEXT)/2, sizeH-1, FOOTER_FG, FOOTER_BG, fmt.Sprintf(FOOTER_TEXT, (*m).Rounds, (*m).Births, (*m).Deaths))
}

// Function tbprint draws a string.
func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}
