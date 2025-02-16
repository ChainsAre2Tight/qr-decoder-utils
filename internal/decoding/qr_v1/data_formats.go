package qr_v1

import (
	"log"

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
	log.Println(length)

	return "no data", nil
}
