package qr_v1

import (
	"fmt"
	"log"
	"reflect"

	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/types"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/utils"
)

type QRVer1 struct{}

func (QRVer1) Decode(matrix [][]bool) (string, error) {
	_, mask, err := readModeAndMask(matrix)
	if err != nil {
		return "", err
	}
	log.Println("selected mask is", reflect.TypeOf(mask))

	format, err := readMetadata(matrix, mask)
	if err != nil {
		return "", err
	}
	log.Println("selected format is", reflect.TypeOf(format))

	data, err := format.ReadData(matrix, mask)
	if err != nil {
		return "", err
	}
	return data, nil
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

type modeInterface interface{}

func readModeAndMask(matrix [][]bool) (modeInterface, maskInterface, error) {
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
	mask, ok := masks[modeString]
	if !ok {
		return nil, nil, fmt.Errorf("no mask matches %s", modeString)
	}

	return nil, mask, nil
}

func readMetadata(matrix [][]bool, mask maskInterface) (formatInterface, error) {
	end := len(matrix) - 1
	rawMetadata := []bool{
		atMatrixXORMask(matrix, mask, end, end),
		atMatrixXORMask(matrix, mask, end-1, end),
		atMatrixXORMask(matrix, mask, end, end-1),
		atMatrixXORMask(matrix, mask, end-1, end-1),
	}

	metadataString := utils.BoolSliceToString(rawMetadata)
	format, ok := formats[metadataString]
	if !ok {
		return nil, fmt.Errorf("format %s is unknown or is not implemented", metadataString)
	}
	return format, nil
}
