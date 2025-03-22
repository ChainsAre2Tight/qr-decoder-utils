package qrdecoder

import (
	"fmt"
	"log"
	"reflect"

	bitreader "github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/QR/common/bit_reader"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/QR/common/data_formats"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/QR/common/masks"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/interfaces"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/utils"
)

func DecodeQR(matrix [][]bool, code interfaces.CodeInterface) (string, error) {
	_, mask, err := readMetadata(matrix)
	if err != nil {
		return "", err
	}

	log.Println("Detected mask:", reflect.TypeOf(mask).Name())
	reader := bitreader.NewBitReader(matrix, mask, code.OOB())

	format, err := readFormat(reader)
	if err != nil {
		return "", err
	}
	log.Println("Detected format:", reflect.TypeOf(format).Name())

	data, err := format.ReadData(matrix, mask, reader)
	if err != nil {
		return "", err
	}
	return data, nil
}

func readMetadata(matrix [][]bool) (interfaces.ModeInterface, interfaces.MaskInterface, error) {
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

func readFormat(reader interfaces.BitReaderInterface) (interfaces.FormatInterface, error) {
	rawMetadata := reader.ReadMultiple(4)

	metadataString := utils.BoolSliceToString(rawMetadata)
	format, ok := data_formats.SUPPORTED_FORMATS[metadataString]
	if !ok {
		return nil, fmt.Errorf("format %s is unknown or is not implemented", metadataString)
	}
	return format, nil
}
