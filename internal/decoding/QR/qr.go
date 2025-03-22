package qr

import (
	qrdecoder "github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/QR/common/qr_decoder"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/interfaces"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/types"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/utils"
)

type QR struct {
	Size              int
	AlignmentPatterns [][2]int
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
	for _, position := range q.AlignmentPatterns {
		if !utils.IsSubmatrix(matrix, types.QRCornerSmall, types.NewPoint(position[0], position[1])) {
			return false
		}

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
	for _, position := range o.QR.AlignmentPatterns {
		if x >= position[0] && x <= position[0]+4 && y >= position[1] && y <= position[1]+4 {
			return true
		}
	}

	return false
}

func (q *QR) OOB() interfaces.OutOfBoundsInterface {
	return &oob{QR: q}
}

func (q *QR) Decode(matrix [][]bool) (string, error) {
	return qrdecoder.DecodeQR(matrix, q)
}
