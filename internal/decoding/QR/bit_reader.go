package qr

import (
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/utils"
)

// it is used to determine if a cell on a matrix
// is a data bit or a part of some pattern
// e.g finder pattern or alignment pattern
type outOfBoundsInterface interface {
	SkipCell(x, y int) bool
	SkipColumn(x int) bool
}

// implements outOfBounds interface for a given QR code
type outOfBounds struct {
	QR *QR
}

// determines if a given column is a part of QR code timing pattern
func (o *outOfBounds) SkipColumn(x int) bool {
	return x == 6
}

// returns true if a given position on a matrix
// is a part of some pattern, e.g. finder pattern
func (o *outOfBounds) SkipCell(x, y int) bool {
	// horizontal timing pattern
	if y == 6 {
		return true
	}
	// finder patterns
	if x <= 8 && y <= 8 || x <= 8 && y >= o.QR.Size-8 || x >= o.QR.Size-8 && y <= 8 {
		return true
	}
	// alignment patterns
	for _, positionX := range o.QR.AlignmentPatterns {
		for _, positionY := range o.QR.AlignmentPatterns {
			// skip alignment patterns that coincide with finder patterns
			if !validAlignmentPattern(positionX, positionY, o.QR.Size) {
				continue
			}
			if x >= positionX-2 && x <= positionY+2 && y >= positionX-2 && y <= positionY+2 {
				return true
			}
		}

	}
	// TODO: check for encoding data (qr v7 and above)

	return false
}

// Sequentially reads a QR code. Every method call
// increments its position on a read sequence.
type bitReader struct {
	matrix   [][]bool
	mask     maskInterface
	position int
	sequence [][2]int
}

// initializes a bit reader
func newBitReader(
	matrix [][]bool,
	mask maskInterface,
	oob outOfBoundsInterface,
) *bitReader {
	return &bitReader{
		matrix:   matrix,
		mask:     mask,
		position: 0,
		sequence: generateReadSequence(len(matrix), len(matrix[0]), oob),
	}
}

// Generates a reading sequence to be used by a bit reader.
// Returns a list of coordinates [x, y] where a N-th element
// is the coordinates of the N-th data bit of a QR code
func generateReadSequence(sizeX, sizeY int, oob outOfBoundsInterface) [][2]int {
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

// Reads one bit from a matrix XORed with the selected mask and increments read sequence position
func (r *bitReader) readOne() bool {
	result := AtMatrixXORMask(r.matrix, r.mask, r.sequence[r.position][0], r.sequence[r.position][1])
	r.position++
	return result
}

// reads N bits from a QR code
func (r *bitReader) readMultiple(n int) []bool {
	result := make([]bool, n)
	for i := range result {
		result[i] = r.readOne()
	}
	// fmt.Println(result)
	return result
}

// Reads one Byte from a QR code
func (r *bitReader) readBytes() byte {
	raw := r.readMultiple(8)
	return byte(utils.BoolSliceToDecimal(raw))
}
