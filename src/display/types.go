package display

import (
	"gitlab.com/ollaww/goameoflife/src/matrix"
)

type Displayable interface {
	Init(*matrix.Matrix) error
	Display(*matrix.Matrix) error
	Update(*matrix.Matrix) error
	Clear(*matrix.Matrix) error
}
