package qr_v1

import (
	qrdecoder "github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/common/qr_decoder"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/interfaces"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/types"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/utils"
)

type QRVer1 struct{}

func (QRVer1) OOB() interfaces.OutOfBoundsInterface {
	return oob{}
}

func (QRVer1) Detect(matrix [][]bool) bool {
	if len(matrix) != 21 {
		return false
	}

	if utils.IsSubmatrix(matrix, types.QRCorner, types.NewPoint(0, 0)) &&
		utils.IsSubmatrix(matrix, types.QRCorner, types.NewPoint(14, 0)) &&
		utils.IsSubmatrix(matrix, types.QRCorner, types.NewPoint(0, 14)) {
		return true
	}

	return false
}

func (q QRVer1) Decode(matrix [][]bool) (string, error) {
	return qrdecoder.DecodeQR(matrix, q)
}
