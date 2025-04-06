package datamatrix

import (
	"fmt"
	"log"
)

func (datamatrix *datamatrix) matrixToBitStream(matrix [][]bool) []bool {

	// convert matrix to array []bool of length datamatrix.X * datamatrix.Y row by row
	array := make([]bool, datamatrix.X*datamatrix.Y)
	for row := range datamatrix.Y {
		for col := range datamatrix.X {
			array[row*datamatrix.X+col] = matrix[col+1][row+1]
		}
	}

	nrow := datamatrix.Y
	ncol := datamatrix.X

	result := make([]bool, datamatrix.X*datamatrix.Y)
	sequence := make([]int, datamatrix.X*datamatrix.Y)
	lock := make([]bool, datamatrix.X*datamatrix.Y)

	module := func(row, col, chr, bit int) {
		if row < 0 {
			row += nrow
			col += 4 - (nrow+4)%8
		}
		if col < 0 {
			col += ncol
			row += 4 - (ncol+4)%8
		}
		result[chr*8+bit-1] = array[row*ncol+col]
		sequence[row*ncol+col] = 10*(chr) + bit
		lock[row*ncol+col] = true
	}

	utah := func(row, col, chr int) {
		module(row-2, col-2, chr, 1)
		module(row-2, col-1, chr, 2)
		module(row-1, col-2, chr, 3)
		module(row-1, col-1, chr, 4)
		module(row-1, col, chr, 5)
		module(row, col-2, chr, 6)
		module(row, col-1, chr, 7)
		module(row, col, chr, 8)
	}

	corner1 := func(chr int) {
		module(nrow-1, 0, chr, 1)
		module(nrow-1, 1, chr, 2)
		module(nrow-1, 2, chr, 3)
		module(0, ncol-2, chr, 4)
		module(0, ncol-1, chr, 5)
		module(1, ncol-1, chr, 6)
		module(2, ncol-1, chr, 7)
		module(3, ncol-1, chr, 8)
	}

	corner2 := func(chr int) {
		module(nrow-3, 0, chr, 1)
		module(nrow-2, 0, chr, 2)
		module(nrow-1, 0, chr, 3)
		module(0, ncol-4, chr, 4)
		module(0, ncol-3, chr, 5)
		module(0, ncol-2, chr, 6)
		module(0, ncol-1, chr, 7)
		module(1, ncol-1, chr, 8)
	}

	corner3 := func(chr int) {
		module(nrow-3, 0, chr, 1)
		module(nrow-2, 0, chr, 2)
		module(nrow-1, 0, chr, 3)
		module(0, ncol-2, chr, 4)
		module(0, ncol-1, chr, 5)
		module(1, ncol-1, chr, 6)
		module(2, ncol-1, chr, 7)
		module(3, ncol-1, chr, 8)
	}

	corner4 := func(chr int) {
		module(nrow-1, 0, chr, 1)
		module(nrow-1, ncol-1, chr, 2)
		module(0, ncol-3, chr, 3)
		module(0, ncol-2, chr, 4)
		module(0, ncol-1, chr, 5)
		module(1, ncol-3, chr, 6)
		module(1, ncol-2, chr, 7)
		module(1, ncol-1, chr, 8)
	}

	chr := 0
	row := 4
	col := 0

	for {
		if row == nrow && col == 0 {
			corner1(chr)
			chr++
		}
		if row == nrow-2 && col == 0 && ncol%4 != 0 {
			corner2(chr)
			chr++
		}
		if row == nrow-2 && col == 0 && ncol%8 == 4 {
			corner3(chr)
			chr++
		}
		if row == nrow+4 && col == 2 && ncol%8 == 0 {
			corner4(chr)
			chr++
		}

		for {
			if row < nrow && col >= 0 && !lock[row*ncol+col] {
				utah(row, col, chr)
				chr++
			}
			row -= 2
			col += 2
			if row < 0 || col >= ncol {
				break
			}
		}

		row += 1
		col += 3

		for {
			if row >= 0 && col <= ncol && !lock[row*ncol+col] {
				utah(row, col, chr)
				chr++
			}
			row += 2
			col -= 2
			if row >= nrow || col < 0 {
				break
			}
		}

		row += 3
		col += 1

		if row >= nrow && col >= ncol {
			break
		}
	}

	// log.Println("lock")
	// printarray(lock, datamatrix.X)
	// log.Println("array")
	// printarray(array, datamatrix.X)
	// log.Println("result")
	// printarray(result, 8)
	log.Println("Reading sequence is:")
	printarray2(sequence, datamatrix.X)

	return result
}

// func printarray(a []bool, n int) {
// 	for row := range len(a) / n {
// 		for col := range n {
// 			if a[row*n+col] {
// 				fmt.Print("1 ")
// 			} else {
// 				fmt.Print("0 ")
// 			}
// 		}
// 		fmt.Print("\n")
// 	}
// }

func printarray2(a []int, n int) {
	for row := range len(a) / n {
		for col := range n {
			fmt.Printf("%5d", a[row*n+col])
		}
		fmt.Print("\n")
	}
}
