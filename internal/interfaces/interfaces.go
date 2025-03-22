package interfaces

import (
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/types"
)

type OutOfBoundsInterface interface {
	SkipCell(x, y int) bool
	SkipColumn(x int) bool
}

type MaskInterface interface {
	At(types.Point) bool
}

type ModeInterface interface{}

type CodeInterface interface {
	Detect([][]bool) bool
	Decode([][]bool) (string, error)
	Description() string
}

type BitReaderInterface interface {
	ReadMultiple(int) []bool
	ReadBytes() byte
}
