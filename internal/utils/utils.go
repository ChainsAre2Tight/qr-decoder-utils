package utils

import (
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
