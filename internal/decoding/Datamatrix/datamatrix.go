package datamatrix

import (
	"fmt"

	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/utils"
)

type datamatrix struct {
	X, Y int
}

func NewDatamatrix(X, Y int) *datamatrix {
	return &datamatrix{
		X: X,
		Y: Y,
	}
}

// Decode implements interfaces.CodeInterface.
func (d *datamatrix) Decode(matrix [][]bool) (string, error) {
	stream := d.matrixToBitStream(matrix)
	fmt.Println(stream)

	for i := range len(stream) / 8 {
		chr := byte(utils.BoolSliceToDecimal(stream[8*i:8*i+1]) - 1)

		fmt.Println(chr)
	}
	return "", nil
}

// Description implements interfaces.CodeInterface.
func (d *datamatrix) Description() string {
	return fmt.Sprintf("ECC 200 Datamatrix (%dx%d)", d.X, d.Y)
}

// Detect implements interfaces.CodeInterface.
func (d *datamatrix) Detect([][]bool) bool {
	panic("unimplemented")
}

func DetetcDatamatrix(matrix [][]bool) (*datamatrix, bool) {
	// check dimensions are even
	if len(matrix)%2 != 0 || len(matrix[0])%2 != 0 {
		return nil, false
	}

	// check L shape
	for _, v := range matrix {
		if !v[len(v)-1] {
			return nil, false
		}
	}
	for _, v := range matrix[0] {
		if !v {
			return nil, false
		}
	}

	// determine size and check timing pattern
	// or just YOLO it

	// construct datamatrix struc based on calculated parameters
	code := NewDatamatrix(len(matrix)-2, len(matrix[0])-2)
	return code, true
}
