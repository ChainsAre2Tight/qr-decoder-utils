package qr

import (
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/types"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/utils"
)

// Checks if any of implemented QR codes
// is present in the matrix
func DetectQR(matrix [][]bool) (*QR, bool) {
	for _, code := range QR_CODES {
		ok := code.Detect(matrix)
		if ok {
			return code, true
		}
	}
	return nil, false
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
