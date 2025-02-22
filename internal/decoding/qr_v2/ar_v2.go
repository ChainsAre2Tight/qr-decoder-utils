package qr_v2

import (
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/qr_v1"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/interfaces"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/types"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/utils"
)

type QRVer2 struct{}

func (QRVer2) Detect(matrix [][]bool) bool {
	if len(matrix) != 25 {
		return false
	}

	if utils.IsSubmatrix(matrix, types.QRCorner, types.NewPoint(0, 0)) &&
		utils.IsSubmatrix(matrix, types.QRCorner, types.NewPoint(18, 0)) &&
		utils.IsSubmatrix(matrix, types.QRCorner, types.NewPoint(0, 18)) &&
		utils.IsSubmatrix(matrix, types.QRCornerSmall, types.NewPoint(16, 16)) {
		return true
	}

	return false
}

func (QRVer2) OOB() interfaces.OutOfBoundsInterface {
	return oob{}
}

// Inherited from QRVer1
func (QRVer2) ReadMetadata(matrix [][]bool) (interfaces.ModeInterface, interfaces.MaskInterface, error) {
	return qr_v1.QRVer1{}.ReadMetadata(matrix)
}

// Inherited from QRVer1
func (QRVer2) ReadFormat(matrix [][]bool, mask interfaces.MaskInterface, reader interfaces.BitReaderInterface) (interfaces.FormatInterface, error) {
	return qr_v1.QRVer1{}.ReadFormat(matrix, mask, reader)
}

type oob struct{}

func (oob) SkipCell(x, y int) bool {
	return y == 6 ||
		x <= 8 && y <= 8 ||
		x <= 8 && y >= 17 ||
		x >= 17 && y <= 8 ||
		x >= 16 && x <= 20 && y >= 16 && y <= 20
}

func (oob) SkipColumn(x int) bool {
	return x == 6
}
