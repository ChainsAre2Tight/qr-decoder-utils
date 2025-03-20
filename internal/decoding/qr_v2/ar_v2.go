package qr_v2

import (
	qrdecoder "github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/common/qr_decoder"
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

func (q QRVer2) Decode(matrix [][]bool) (string, error) {
	return qrdecoder.DecodeQR(matrix, q)
}
