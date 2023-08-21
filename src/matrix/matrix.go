package matrix

import (
	"math/rand"
)

func Create(rows, cols int, p float64) *Matrix {
	m := make([][]bool, rows)
	for k := range m {
		m[k] = make([]bool, cols)
		for j := range m[k] {
			m[k][j] = (rand.Float64() > (1 - p))
		}
	}
	return &Matrix{
		Matrix: &m,
	}
}

func (current *Matrix) ComputeNextOn(next *Matrix) {
	currentMatrix := *current.Matrix
	nextMatrix := *next.Matrix
	for k := range currentMatrix {
		for j := range (currentMatrix)[k] {
			aliveNeighbours := current.aliveNeighbours(k, j)
			if currentMatrix[k][j] {
				nextMatrix[k][j] = aliveNeighbours == 2 || aliveNeighbours == 3
				if !nextMatrix[k][j] {
					current.Deaths++
				}
			} else {
				nextMatrix[k][j] = aliveNeighbours == 3
				if nextMatrix[k][j] {
					current.Births++ // A star is born
				}
			}
		}
	}
	next.Rounds = current.Rounds + 1
	next.Births = current.Births
	next.Deaths = current.Deaths
}

func (m *Matrix) aliveNeighbours(x, y int) int {
	alive := 0
	_matrix := *m.Matrix
	offests := [3]int{-1, 0, 1}
	for _, xo := range offests {
		for _, yo := range offests {
			if m.areValidCoordinatates(x+xo, y+yo) && _matrix[x+xo][y+yo] {
				alive++
			}
		}
	}
	return alive
}

func (m *Matrix) areValidCoordinatates(x, y int) bool {
	row, col := m.dimensions()
	return x >= 0 && y >= 0 && x < row && y < col
}

func (m *Matrix) dimensions() (int, int) {
	mtr := *m.Matrix
	return len(mtr), len((mtr)[0])
}
