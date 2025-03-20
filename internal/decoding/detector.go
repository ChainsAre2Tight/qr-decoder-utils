package decoding

import (
	"fmt"

	qrdetection "github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/QR/common/type_detection"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/interfaces"
)

func detectCodeType(matrix [][]bool) (interfaces.CodeInterface, error) {

	code, ok := qrdetection.DetectQR(matrix)
	if ok {
		return code, nil
	}
	return nil, fmt.Errorf("no known code types found")
}
