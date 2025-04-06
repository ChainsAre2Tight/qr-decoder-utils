package decoding

import (
	"fmt"

	datamatrix "github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/Datamatrix"
	qr "github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/QR"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/interfaces"
)

func detectCodeType(matrix [][]bool) (interfaces.CodeInterface, error) {

	if code, ok := qr.DetectQR(matrix); ok {
		return code, nil
	}

	if code, ok := datamatrix.DetectDatamatrix(matrix); ok {
		return code, nil
	}
	return nil, fmt.Errorf("no known code types found")
}
