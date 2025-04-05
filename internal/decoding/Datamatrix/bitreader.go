package datamatrix

import "fmt"

func (datamatrix *datamatrix) matrixToBitStream(matrix [][]bool) []bool {
	result := make([]bool, datamatrix.X*datamatrix.Y-4)

	// starting in the correct location for character #1
	chr := 1
	row := 4
	col := 0

	nrow := datamatrix.Y
	ncol := datamatrix.X

	// repeatedly first check for one of the special corner cases, then
	for {
		if row == nrow && col == 0 {
			readCorner1(ncol, nrow, chr, result, matrix)
			chr++
		}
		if row == nrow-2 && col == 0 && ncol%4 != 0 {
			readCorner2(ncol, nrow, chr, result, matrix)
			chr++
		}
		if row == nrow-2 && col == 0 && ncol%8 == 4 {
			readCorner3(ncol, nrow, chr, result, matrix)
			chr++
		}
		if row == nrow+4 && col == 2 && ncol%8 == 0 {
			readCorner4(ncol, nrow, chr, result, matrix)
		}

		// sweep upward diagonnally, insering successive characters
		for {
			if row < nrow && col >= 0 && !matrix[col][row] {
				readUtah(row, col, chr, result, matrix)
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

		// & then sweep downward diagonnally, inserting successive haracters XDD
		for {
			if row >= 0 && col <= ncol && !matrix[col][row] {
				readUtah(row, col, chr, result, matrix)
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

		// until the entire matrix is scanned
		if row >= nrow && col >= ncol {
			break
		}
	}

	// lastly, if the lower righthand corner is untouched, fill in fixed pattern
	// if !matrix[ncol-1][nrow] {
	// 	result[nrow*ncol-1] = true
	// 	result[nrow*ncol-ncol-2] = true
	// }

	return result
}

func at(matrix [][]bool, row, col int) bool {
	return matrix[col+1][row+1]
}

func readUtah(row, col, chr int, result []bool, matrix [][]bool) {
	fmt.Println("utah", row, col, chr)
	result[chr] = at(matrix, row-2, col-2)
	result[chr+1] = at(matrix, row-2, col-1)
	result[chr+2] = at(matrix, row-1, col-2)
	result[chr+3] = at(matrix, row-1, col-1)
	result[chr+4] = at(matrix, row-1, col)
	result[chr+5] = at(matrix, row, col-2)
	result[chr+6] = at(matrix, row, col-1)
	result[chr+7] = at(matrix, row, col)
}

func readCorner1(ncol, nrow, chr int, result []bool, matrix [][]bool) {
	fmt.Println("corner1", chr)
	result[chr] = at(matrix, nrow-1, 0)
	result[chr+1] = at(matrix, nrow-1, 1)
	result[chr+2] = at(matrix, nrow-1, 2)
	result[chr+3] = at(matrix, 0, ncol-2)
	result[chr+4] = at(matrix, 0, ncol-1)
	result[chr+5] = at(matrix, 1, ncol-1)
	result[chr+6] = at(matrix, 2, ncol-1)
	result[chr+6] = at(matrix, 3, ncol-1)
}

func readCorner2(ncol, nrow, chr int, result []bool, matrix [][]bool) {
	fmt.Println("corner2", chr)
	result[chr] = at(matrix, nrow-3, 0)
	result[chr+1] = at(matrix, nrow-2, 0)
	result[chr+2] = at(matrix, nrow-1, 0)
	result[chr+3] = at(matrix, 0, ncol-4)
	result[chr+4] = at(matrix, 0, ncol-3)
	result[chr+5] = at(matrix, 0, ncol-2)
	result[chr+6] = at(matrix, 0, ncol-1)
	result[chr+6] = at(matrix, 1, ncol-1)
}

func readCorner3(ncol, nrow, chr int, result []bool, matrix [][]bool) {
	fmt.Println("corner3", chr)
	result[chr] = at(matrix, nrow-3, 0)
	result[chr+1] = at(matrix, nrow-2, 0)
	result[chr+2] = at(matrix, nrow-1, 0)
	result[chr+3] = at(matrix, 0, ncol-2)
	result[chr+4] = at(matrix, 0, ncol-1)
	result[chr+5] = at(matrix, 1, ncol-1)
	result[chr+6] = at(matrix, 2, ncol-1)
	result[chr+6] = at(matrix, 3, ncol-1)
}

func readCorner4(ncol, nrow, chr int, result []bool, matrix [][]bool) {
	fmt.Println("corner4", chr)
	result[chr] = at(matrix, nrow-1, 0)
	result[chr+1] = at(matrix, nrow-1, ncol-1)
	result[chr+2] = at(matrix, 0, ncol-3)
	result[chr+3] = at(matrix, 0, ncol-2)
	result[chr+4] = at(matrix, 0, ncol-1)
	result[chr+5] = at(matrix, 1, ncol-3)
	result[chr+6] = at(matrix, 1, ncol-2)
	result[chr+6] = at(matrix, 1, ncol-1)
}
