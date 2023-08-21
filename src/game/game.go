package game

import (
	"time"

	"gitlab.com/ollaww/goameoflife/src/display"
	"gitlab.com/ollaww/goameoflife/src/matrix"
)

const TIME_INTERVAL time.Duration = 1000

type GoameOfLife struct {
	currentMatrix *matrix.Matrix
	nextMatrix    *matrix.Matrix
	displayer     display.Displayable
}

func CreateInstance() *GoameOfLife {
	instance := GoameOfLife{
		currentMatrix: matrix.Create(
			matrix.DEFAULT_MATRIX_ROWS,
			matrix.DEFAULT_MATRIX_COLS,
			matrix.DEFAULT_PROBABILITY,
		),
		nextMatrix: matrix.Create(
			matrix.DEFAULT_MATRIX_ROWS,
			matrix.DEFAULT_MATRIX_COLS,
			0),
		displayer: &display.TermboxDisplayer{},
	}
	return &instance
}

func (g *GoameOfLife) Start() {
	displayer := g.displayer
	g.init()
	displayer.Display(g.currentMatrix)
	for {
		time.Sleep(TIME_INTERVAL * time.Millisecond)
		g.nextRound()
		g.displayer.Update(g.currentMatrix)
	}
}

func (g *GoameOfLife) init() {
	displayError := g.displayer.Init(g.currentMatrix)
	if displayError != nil {
		panic(displayError)
	}
}

func (g *GoameOfLife) nextRound() {
	current := g.currentMatrix
	next := g.nextMatrix
	holder := current

	current.ComputeNextOn(next)
	g.currentMatrix = next
	g.nextMatrix = holder
}
