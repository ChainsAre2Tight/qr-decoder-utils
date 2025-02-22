package qr_v2

import (
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
	panic("not implemented")
}

func (QRVer2) ReadMetadata(matrix [][]bool) (interfaces.ModeInterface, interfaces.MaskInterface, error) {
	panic("not implemented")
}

func (QRVer2) ReadFormat(matrix [][]bool, _ interfaces.MaskInterface, _ interfaces.BitReaderInterface) (interfaces.FormatInterface, error) {
	panic("not implemented")
}
