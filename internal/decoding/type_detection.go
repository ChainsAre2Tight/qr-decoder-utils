package decoding

import (
	"fmt"

	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/qr_v1"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/types"
)

var KNOWN_CODES = []types.CodeInterface{
	qr_v1.QRVer1{},
}

func DetectCodeType(matrix [][]bool) (types.CodeInterface, error) {

	for _, code := range KNOWN_CODES {
		ok := code.Detect(matrix)
		if ok {
			return code, nil
		}
	}
	return nil, fmt.Errorf("no known code types found")
}
