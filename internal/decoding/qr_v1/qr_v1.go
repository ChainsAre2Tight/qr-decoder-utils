package qr_v1

import (
	"fmt"

	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/common/data_formats"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/common/masks"
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

func (QRVer1) ReadMetadata(matrix [][]bool) (interfaces.ModeInterface, interfaces.MaskInterface, error) {
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

func (QRVer1) ReadFormat(matrix [][]bool, mask interfaces.MaskInterface, reader interfaces.BitReaderInterface) (interfaces.FormatInterface, error) {
	rawMetadata := reader.ReadMultiple(4)

	metadataString := utils.BoolSliceToString(rawMetadata)
	format, ok := data_formats.SUPPORTED_FORMATS[metadataString]
	if !ok {
		return nil, fmt.Errorf("format %s is unknown or is not implemented", metadataString)
	}
	return format, nil
}

func (q QRVer1) Decode(matrix [][]bool) (string, error) {
	return qrdecoder.DecodeQR(matrix, q)
}
