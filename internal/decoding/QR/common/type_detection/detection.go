package type_detection

import (
	qr "github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/QR"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/interfaces"
)

// All hardcoded qr codes used in this app
// Refer to table 1 for size parameter and
// table E.1 for list of alignment patterns
var QR_CODES = []interfaces.CodeInterface{
	&qr.QR{Name: "QR Version 1", Size: 21, AlignmentPatterns: []int{}},
	&qr.QR{Name: "QR Version 2", Size: 25, AlignmentPatterns: []int{6, 18}},
	&qr.QR{Name: "QR Version 3", Size: 29, AlignmentPatterns: []int{6, 22}},
	&qr.QR{Name: "QR Version 4", Size: 33, AlignmentPatterns: []int{6, 26}},
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
