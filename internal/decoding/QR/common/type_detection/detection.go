package type_detection

import (
	qr "github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/QR"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/interfaces"
)

var QR_CODES = []interfaces.CodeInterface{
	&qr.QR{Size: 21},
	&qr.QR{Size: 25, AlignmentPatterns: [][2]int{{16, 16}}},
	&qr.QR{Size: 29, AlignmentPatterns: [][2]int{{20, 20}}},
	&qr.QR{Size: 33, AlignmentPatterns: [][2]int{{24, 24}}},
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
