package matrix

const DEFAULT_MATRIX_ROWS int = 10
const DEFAULT_MATRIX_COLS int = 25
const DEFAULT_PROBABILITY float64 = .55

type Matrix struct {
	Matrix *[][]bool
	Rounds uint
	Births uint
	Deaths uint
}
