package qr_v2

import (
	"fmt"

	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/common/data_formats"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/common/masks"
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

func (QRVer2) ReadMetadata(matrix [][]bool) (interfaces.ModeInterface, interfaces.MaskInterface, error) {
	// omit first two bits, mode is not implemented
	mode, err := utils.ReadMatrixRow(matrix, 8, 2, 5)
	if err != nil {
		return nil, nil, err
	}

	mode, err = utils.XORSlices(mode, []bool{true, false, true})
	if err != nil {
		return nil, nil, err
	}

	modeString := utils.BoolSliceToString(mode)
	mask, ok := masks.Masks[modeString]
	if !ok {
		return nil, nil, fmt.Errorf("no mask matches %s", modeString)
	}

	return nil, mask, nil
}

func (QRVer2) ReadFormat(matrix [][]bool, mask interfaces.MaskInterface, reader interfaces.BitReaderInterface) (interfaces.FormatInterface, error) {
	rawMetadata := reader.ReadMultiple(4)

	metadataString := utils.BoolSliceToString(rawMetadata)
	format, ok := data_formats.SUPPORTED_FORMATS[metadataString]
	if !ok {
		return nil, fmt.Errorf("format %s is unknown or is not implemented", metadataString)
	}
	return format, nil
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
