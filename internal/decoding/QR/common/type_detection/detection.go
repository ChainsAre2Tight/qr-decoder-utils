package type_detection

import (
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/QR/qr_v1"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/QR/qr_v2"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/QR/qr_v3"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/interfaces"
)

var QR_CODES = []interfaces.CodeInterface{
	qr_v1.QRVer1{},
	qr_v2.QRVer2{},
	qr_v3.QRVer3{},
}

func DetectQR(matrix [][]bool) (interfaces.CodeInterface, bool) {
	for _, code := range QR_CODES {
		ok := code.Detect(matrix)
		if ok {
			return code, true
		}
	}
	return nil, false
}
