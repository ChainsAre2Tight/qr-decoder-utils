package decoding

import (
	"fmt"

	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/qr_v1"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/qr_v2"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/qr_v3"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/interfaces"
)

var KNOWN_CODES = []interfaces.CodeInterface{
	qr_v1.QRVer1{},
	qr_v2.QRVer2{},
	qr_v3.QRVer3{},
}

func detectCodeType(matrix [][]bool) (interfaces.CodeInterface, error) {

	for _, code := range KNOWN_CODES {
		ok := code.Detect(matrix)
		if ok {
			return code, nil
		}
	}
	return nil, fmt.Errorf("no known code types found")
}
