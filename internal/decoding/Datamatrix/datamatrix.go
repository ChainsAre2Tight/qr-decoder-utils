package datamatrix

import "fmt"

type datamatrix struct {
	x, y int
}

func NewDatamatrix(X, Y int) *datamatrix {
	return &datamatrix{
		x: X,
		y: Y,
	}
}

// Decode implements interfaces.CodeInterface.
func (d *datamatrix) Decode([][]bool) (string, error) {
	panic("unimplemented")
}

// Description implements interfaces.CodeInterface.
func (d *datamatrix) Description() string {
	return fmt.Sprintf("ECC 200 Datamatrix (%dx%d)", d.x, d.y)
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
