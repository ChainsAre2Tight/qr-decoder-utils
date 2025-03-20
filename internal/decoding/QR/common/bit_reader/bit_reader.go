package bitreader

import (
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/QR/common/masks"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/interfaces"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/utils"
)

type bitReader struct {
	matrix   [][]bool
	mask     interfaces.MaskInterface
	position int
	sequence [][2]int
}

func NewBitReader(
	matrix [][]bool,
	mask interfaces.MaskInterface,
	oob interfaces.OutOfBoundsInterface,
) *bitReader {
	return &bitReader{
		matrix:   matrix,
		mask:     mask,
		position: 0,
		sequence: GenerateReadSequence(len(matrix), len(matrix[0]), oob),
	}
}

// !!! megagovnokod !!! N^2 time, N^2 space
func GenerateReadSequence(sizeX, sizeY int, oob interfaces.OutOfBoundsInterface) [][2]int {
	up := true

	result := make([][2]int, 0, sizeX*sizeY)

	// iterate throuth 2-wide columns in reverse order
	for mainX := sizeX - 1; mainX >= 0; mainX -= 2 {
		if oob.SkipColumn(mainX) {
			mainX--
		}

		// main counter for Y coordinate, ignores direction
		for mainY := range sizeY {
			var y int
			// if going up, Y is actually sizeY - mainY - 1 (reverse order)
			if up {
				y = sizeY - mainY - 1
			} else {
				y = mainY
			}

			for x := mainX; x >= mainX-1; x-- {
				if x < 0 || y < 0 || oob.SkipCell(x, y) {
					continue
				}
				result = append(result, [2]int{x, y})
			}
		}
		up = !up
	}
	return result
}

func (r *bitReader) readOne() bool {
	result := masks.AtMatrixXORMask(r.matrix, r.mask, r.sequence[r.position][0], r.sequence[r.position][1])
	r.position++
	return result
}

func (r *bitReader) ReadMultiple(n int) []bool {
	result := make([]bool, n)
	for i := range result {
		result[i] = r.readOne()
	}
	// fmt.Println(result)
	return result
}

func (r *bitReader) ReadBytes() byte {
	raw := r.ReadMultiple(8)
	return byte(utils.BoolSliceToDecimal(raw))
}
