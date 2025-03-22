package data_formats

import (
	"fmt"
	"log"
	"strings"

	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/interfaces"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/utils"
)

type integerFormat struct{}

type byteFormat struct{}

var SUPPORTED_FORMATS = map[string]interfaces.FormatInterface{
	"0001": integerFormat{},
	"0100": byteFormat{},
}

func (byteFormat) ReadData(matrix [][]bool, mask interfaces.MaskInterface, reader interfaces.BitReaderInterface) (string, error) {
	length := utils.BoolSliceToDecimal(reader.ReadMultiple(8))
	log.Println("Detected content length:", length)

	raw := make([]byte, length)

	for i := range length {
		raw[i] = reader.ReadBytes()
	}

	data, err := utils.BytesToISO8859dash1(raw)

	return data, err
}

func (integerFormat) ReadData(matrix [][]bool, mask interfaces.MaskInterface, reader interfaces.BitReaderInterface) (string, error) {
	// read length
	length := utils.BoolSliceToDecimal(reader.ReadMultiple(10))
	log.Println("Detected content length:", length)

	var resultBuilder strings.Builder

	// count 3-groups
	countFull := length / 3

	// read 3-groups
	for range countFull {
		raw := utils.BoolSliceToDecimal(reader.ReadMultiple(10))
		resultBuilder.WriteString(fmt.Sprintf("%03d", raw))
	}

	countRemainder := length % 3
	var raw int

	if countRemainder == 1 {
		raw = utils.BoolSliceToDecimal(reader.ReadMultiple(4))
		resultBuilder.WriteString(fmt.Sprintf("%01d", raw))
	} else {
		raw = utils.BoolSliceToDecimal(reader.ReadMultiple(7))
		resultBuilder.WriteString(fmt.Sprintf("%02d", raw))
	}

	return resultBuilder.String(), nil
}
