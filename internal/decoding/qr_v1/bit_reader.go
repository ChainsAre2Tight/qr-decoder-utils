package qr_v1

import (
	"fmt"
	"log"
)

type bitReader struct {
	matrix        [][]bool
	mask          maskInterface
	x, y          int
	flagShift     int
	flagDirection int
	counter       int
}

func NewBitReader(matrix [][]bool, mask maskInterface) bitReader {
	return bitReader{
		matrix:        matrix,
		mask:          mask,
		x:             len(matrix) - 1,
		y:             len(matrix[0]) - 3,
		flagShift:     -1,
		flagDirection: -1,
		counter:       2,
	}
}

func (r *bitReader) ReadOne() (bool, int, int) {
	value := atMatrixXORMask(r.matrix, r.mask, r.x, r.y)
	resultX, resultY := r.x, r.y
	log.Printf("at %d, %d: %t", r.x, r.y, value)

	r.calcNextPos()
	return value, resultX, resultY
}

func (r *bitReader) ReadMultiple(n int) []bool {
	result := make([]bool, n)
	for i := range result {
		result[i], _, _ = r.ReadOne()
	}
	fmt.Println(result)
	return result
}

func (r *bitReader) calcNextPos() {

	if r.counter == 24 {
		r.flagShift = -1
		r.flagDirection *= -1
	}

	r.x = r.x + r.flagShift
	r.y += r.flagDirection * ((2 - r.flagShift) % 3)

	r.flagShift *= -1

	if r.counter == 24 {
		r.counter = 0
		r.flagShift = -1
	}

	r.counter++
}
