package qr_v1

import (
	"fmt"
	"log"
	"strings"

	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/utils"
)

type formatInterface interface {
	ReadData([][]bool, maskInterface) (string, error)
}

type integerFormat struct{}

// type byteFormat struct{}

var formats = map[string]formatInterface{
	"0001": integerFormat{},
	// "0100": byteFormat{},
}

func (integerFormat) ReadData(matrix [][]bool, mask maskInterface) (string, error) {
	reader := NewBitReader(matrix, mask)
	// read length
	length := utils.BoolSliceToDecimal(reader.ReadMultiple(10))
	log.Println("detected content length:", length)

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
