package decoding

import (
	"fmt"

	qr "github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/QR"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/interfaces"
)

func detectCodeType(matrix [][]bool) (interfaces.CodeInterface, error) {

	code, ok := qr.DetectQR(matrix)
	if ok {
		return code, nil
	}
	return nil, fmt.Errorf("no known code types found")
}
