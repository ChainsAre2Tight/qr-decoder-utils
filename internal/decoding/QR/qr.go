package qr

import (
	"fmt"

	qrdecoder "github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/QR/common/qr_decoder"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/interfaces"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/types"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/utils"
)

var CCI1dash9 = &types.CCI{Numeric: 10, Alphanumeric: 9, Byte: 16, Kanji: 8}
var CCI10dash26 = &types.CCI{Numeric: 12, Alphanumeric: 11, Byte: 16, Kanji: 10}
var CCI27dash40 = &types.CCI{Numeric: 14, Alphanumeric: 13, Byte: 16, Kanji: 12}

type QR struct {
	Name              string
	Size              int
	Cci               *types.CCI
	AlignmentPatterns []int
}

// Performs checks on a given matrix to determine if it contains
// a QR code of specified parameters
func (q *QR) Detect(matrix [][]bool) bool {

	// check basic dimensions
	if len(matrix) != q.Size || len(matrix[0]) != q.Size {
		return false
	}

	// check for finder patterns
	if !utils.IsSubmatrix(matrix, types.QRCorner, types.NewPoint(0, 0)) ||
		!utils.IsSubmatrix(matrix, types.QRCorner, types.NewPoint(q.Size-7, 0)) ||
		!utils.IsSubmatrix(matrix, types.QRCorner, types.NewPoint(0, q.Size-7)) {
		return false
	}

	// check for alignment patterns, if any
	for _, positionX := range q.AlignmentPatterns {
		for _, positionY := range q.AlignmentPatterns {

			// skip alignment patterns that coincide with finder patterns
			if !validAlignmentPattern(positionX, positionY, q.Size) {
				continue
			}

			if !utils.IsSubmatrix(matrix, types.QRCornerSmall, types.NewPoint(positionX-2, positionY-2)) {
				return false
			}
		}
	}

	return true
}

// checks if an alignment pattern coincides with a finder pattern
func validAlignmentPattern(centerX, centerY, size int) bool {
	if centerX == 6 && centerY == 6 || centerX == size-7 && centerY == 6 || centerX == 6 && centerY == size-7 {
		return false
	}
	return true
}

type oob struct {
	QR *QR
}

func (o *oob) SkipColumn(x int) bool {
	return x == 6
}

func (o *oob) SkipCell(x, y int) bool {
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
	// TODO: check for encoding data

	return false
}

func (q *QR) OOB() interfaces.OutOfBoundsInterface {
	return &oob{QR: q}
}

func (q *QR) Decode(matrix [][]bool) (string, error) {
	return qrdecoder.DecodeQR(matrix, q)
}

func (q *QR) Description() string {
	return fmt.Sprintf("%s (%dx%d)", q.Name, q.Size, q.Size)
}
