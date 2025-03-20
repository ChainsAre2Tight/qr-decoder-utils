package qr_v3

import (
	qrdecoder "github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/common/qr_decoder"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/qr_v1"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/interfaces"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/types"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/utils"
)

type QRVer3 struct{}

func (QRVer3) Detect(matrix [][]bool) bool {
	if len(matrix) != 29 {
		return false
	}

	if utils.IsSubmatrix(matrix, types.QRCorner, types.NewPoint(0, 0)) &&
		utils.IsSubmatrix(matrix, types.QRCorner, types.NewPoint(22, 0)) &&
		utils.IsSubmatrix(matrix, types.QRCorner, types.NewPoint(0, 22)) &&
		utils.IsSubmatrix(matrix, types.QRCornerSmall, types.NewPoint(20, 20)) {
		return true
	}

	return false
}

func (QRVer3) OOB() interfaces.OutOfBoundsInterface {
	return oob{}
}

// Inherited from QRVer1
func (QRVer3) ReadMetadata(matrix [][]bool) (interfaces.ModeInterface, interfaces.MaskInterface, error) {
	return qr_v1.QRVer1{}.ReadMetadata(matrix)
}

// Inherited from QRVer1
func (QRVer3) ReadFormat(matrix [][]bool, mask interfaces.MaskInterface, reader interfaces.BitReaderInterface) (interfaces.FormatInterface, error) {
	return qr_v1.QRVer1{}.ReadFormat(matrix, mask, reader)
}

type oob struct{}

func (oob) SkipCell(x, y int) bool {
	return y == 6 ||
		x <= 8 && y <= 8 ||
		x <= 8 && y >= 21 ||
		x >= 21 && y <= 8 ||
		x >= 20 && x <= 24 && y >= 20 && y <= 24
}

func (oob) SkipColumn(x int) bool {
	return x == 6
}

func (q QRVer3) Decode(matrix [][]bool) (string, error) {
	return qrdecoder.DecodeQR(matrix, q)
}
