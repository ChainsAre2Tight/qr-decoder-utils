package utils

import (
	"fmt"
	"math"
	"strings"

	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/types"
)

func Concat(a, b string) string {
	var sb strings.Builder
	sb.WriteString(a)
	sb.WriteString(b)
	return sb.String()
}

func IsSubmatrix(matA, matB [][]bool, start types.Point) bool {
	for x := range matB {
		if x+start.X >= len(matA) {
			return false
		}
		for y := range matB[x] {
			if y+start.Y >= len(matA[x]) || matA[start.X+x][start.Y+y] != matB[x][y] {
				return false
			}
		}
	}
	return true
}

func ReadMatrixRow(matrix [][]bool, row, start, end int) ([]bool, error) {
	if start > end || row >= len(matrix[0]) || end > len(matrix) {
		return []bool{}, fmt.Errorf("invalid start and end positions")
	}
	result := make([]bool, end-start)
	for x := range result {
		result[x] = matrix[start+x][row]
	}
	return result, nil
}

func XORSlices(sliceA, sliceB []bool) ([]bool, error) {
	if len(sliceA) != len(sliceB) {
		return nil, fmt.Errorf("slices must be of equal length")
	}
	result := make([]bool, len(sliceA))
	for i := range sliceA {
		result[i] = sliceA[i] != sliceB[i]
	}
	return result, nil
}

func BoolSliceToString(slice []bool) string {
	var sb strings.Builder
	for _, val := range slice {
		if val {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
	}
	return sb.String()
}

func BoolSliceToDecimal(slice []bool) int {
	result := 0
	for i := range slice {
		if slice[len(slice)-1-i] {
			result += int(math.Pow(float64(2), float64(i)))
		}
	}
	return result
}
