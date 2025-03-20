package qrdecoder

import (
	"log"
	"reflect"

	bitreader "github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/common/bit_reader"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/interfaces"
)

func DecodeQR(matrix [][]bool, code interfaces.CodeInterface) (string, error) {
	_, mask, err := code.ReadMetadata(matrix)
	if err != nil {
		return "", err
	}

	log.Println("Detected mask:", reflect.TypeOf(mask).Name())
	reader := bitreader.NewBitReader(matrix, mask, code.OOB())

	format, err := code.ReadFormat(matrix, mask, reader)
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
