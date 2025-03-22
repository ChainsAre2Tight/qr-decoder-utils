package qr

import (
	"fmt"
	"log"
	"reflect"

	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/interfaces"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/utils"
)

// Decodes a QR code in a given matrix
func (q *QR) Decode(matrix [][]bool) (string, error) {
	_, mask, err := readMetadata(matrix)
	if err != nil {
		return "", err
	}

	log.Println("Detected mask:", reflect.TypeOf(mask))
	reader := newBitReader(matrix, mask, &outOfBounds{QR: q})

	format, err := readFormat(reader)
	if err != nil {
		return "", err
	}
	log.Println("Detected format:", reflect.TypeOf(format))

	data, err := format.ReadData(matrix, mask, reader, q.Cci)
	if err != nil {
		return "", err
	}
	return data, nil
}

// Reads Mode and Mask of a QR coder in a given matrix
func readMetadata(matrix [][]bool) (interfaces.ModeInterface, maskInterface, error) {
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
	mask, ok := Masks[modeString]
	if !ok {
		return nil, nil, fmt.Errorf("no mask matches %s", modeString)
	}

	return nil, mask, nil
}

// reads data fromat of a QR code in a given matrix
func readFormat(reader *bitReader) (formatInterface, error) {
	rawMetadata := reader.readMultiple(4)

	metadataString := utils.BoolSliceToString(rawMetadata)
	format, ok := SUPPORTED_FORMATS[metadataString]
	if !ok {
		return nil, fmt.Errorf("format %s is unknown or is not implemented", metadataString)
	}
	return format, nil
}
