package encoding

import (
	fec "github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/FEC"
	qr "github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/QR"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/types"
)

func Encode(...any) types.Matrix {
	code := qr.QR_CODES[3]

	result := code.Encode([]bool{false, true, false, true, false}, fec.Modes["L"], qr.Mask110{})

	return result
}
