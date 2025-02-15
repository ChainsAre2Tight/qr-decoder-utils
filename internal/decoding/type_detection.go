package decoding

import (
	"fmt"
)

type CodeInterface interface {
	Detect([][]bool) bool
	Decode([][]bool) (string, error)
}

type QRVer1 struct{}

var KNOWN_CODES = []CodeInterface{
	QRVer1{},
}

var QRCorner = [][]bool{
	{true, true, true, true, true, true, true},
	{true, false, false, false, false, false, true},
	{true, false, true, true, true, false, true},
	{true, false, true, true, true, false, true},
	{true, false, true, true, true, false, true},
	{true, false, false, false, false, false, true},
	{true, true, true, true, true, true, true},
}

func DetectCodeType(matrix [][]bool) (CodeInterface, error) {

	for _, code := range KNOWN_CODES {
		ok := code.Detect(matrix)
		if ok {
			return code, nil
		}
	}
	return nil, fmt.Errorf("no known code types found")
}
