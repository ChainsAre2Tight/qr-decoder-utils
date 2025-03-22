package qr

import (
	"fmt"
	"log"
	"strings"

	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/utils"
)

type formatInterface interface {
	// Reads data of a QR code.
	ReadData([][]bool, maskInterface, *bitReader, *cci) (string, error)
}

type integerFormat struct{}
type byteFormat struct{}

var SUPPORTED_FORMATS = map[string]formatInterface{
	"0001": &integerFormat{},
	"0100": &byteFormat{},
}

func (byteFormat) ReadData(matrix [][]bool, mask maskInterface, reader *bitReader, cci *cci) (string, error) {
	length := utils.BoolSliceToDecimal(reader.readMultiple(cci.Byte))

	log.Println("Detected content length:", length)

	raw := make([]byte, length)

	for i := range length {
		raw[i] = reader.readBytes()
	}

	// attempt utf 8 decoding
	data, err := utils.BytesToUTF8(raw)
	if err == nil {
		log.Println("Content is a valid UTF-8 string, decoding...")
		return data, err
	}
	// default to ISO8859-1
	log.Println("Defaulting to ISO8859-1")
	data, err = utils.BytesToISO8859dash1(raw)

	return data, err
}

func (integerFormat) ReadData(matrix [][]bool, mask maskInterface, reader *bitReader, cci *cci) (string, error) {
	// read length
	length := utils.BoolSliceToDecimal(reader.readMultiple(cci.Numeric))

	log.Println("Detected content length:", length)

	var resultBuilder strings.Builder

	// count 3-groups
	countFull := length / 3

	// read 3-groups
	for range countFull {
		raw := utils.BoolSliceToDecimal(reader.readMultiple(10))
		resultBuilder.WriteString(fmt.Sprintf("%03d", raw))
	}

	countRemainder := length % 3
	var raw int

	if countRemainder == 1 {
		raw = utils.BoolSliceToDecimal(reader.readMultiple(4))
		resultBuilder.WriteString(fmt.Sprintf("%01d", raw))
	} else {
		raw = utils.BoolSliceToDecimal(reader.readMultiple(7))
		resultBuilder.WriteString(fmt.Sprintf("%02d", raw))
	}

	return resultBuilder.String(), nil
}
