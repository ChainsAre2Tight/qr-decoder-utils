package decoding

import (
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/types"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/utils"
)

func (QRVer1) Decode(matrix [][]bool) (string, error) {
	return "", nil
}

func (QRVer1) Detect(matrix [][]bool) bool {
	if len(matrix) != 21 {
		return false
	}

	if utils.IsSubmatrix(matrix, QRCorner, types.NewPoint(0, 0)) &&
		utils.IsSubmatrix(matrix, QRCorner, types.NewPoint(14, 0)) &&
		utils.IsSubmatrix(matrix, QRCorner, types.NewPoint(0, 14)) {
		return true
	}

	return false
}
