package qr_v1

import (
	"fmt"
)

var readSequence = [][2]int{
	{20, 18}, {19, 18}, {20, 17}, {19, 17}, {20, 16}, {19, 16}, {20, 15}, {19, 15}, // up
	{20, 14}, {19, 14}, {20, 13}, {19, 13}, {20, 12}, {19, 12}, {20, 11}, {19, 11}, // up
	{20, 10}, {19, 10}, {20, 9}, {19, 9}, {18, 9}, {17, 9}, {18, 10}, {17, 10}, // left
	{18, 11}, {17, 11}, {18, 12}, {17, 12}, {18, 13}, {17, 13}, {18, 14}, {17, 14}, // down
	{18, 15}, {17, 15}, {18, 16}, {17, 16}, {18, 17}, {17, 17}, {18, 18}, {17, 18}, // down
	{18, 19}, {17, 19}, {18, 20}, {17, 20}, {16, 20}, {15, 20}, {16, 19}, {15, 19}, // left

}

type bitReader struct {
	matrix   [][]bool
	mask     maskInterface
	position int
}

func NewBitReader(matrix [][]bool, mask maskInterface) bitReader {
	return bitReader{
		matrix:   matrix,
		mask:     mask,
		position: 0,
	}
}

func (r *bitReader) ReadOne() (bool, int, int) {
	x, y := readSequence[r.position][0], readSequence[r.position][1]
	value := atMatrixXORMask(r.matrix, r.mask, x, y)
	defer func(r *bitReader) { r.position++ }(r) // shift after returning
	return value, x, y
}

func (r *bitReader) ReadMultiple(n int) []bool {
	result := make([]bool, n)
	for i := range result {
		result[i], _, _ = r.ReadOne()
	}
	fmt.Println(result)
	return result
}
