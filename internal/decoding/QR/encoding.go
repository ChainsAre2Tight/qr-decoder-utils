package qr

import (
	fec "github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/FEC"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/types"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/utils"
)

func (q *QR) Encode(data []bool, mode fec.Mode, mask maskInterface) types.Matrix {

	// initialize result
	matrix := make(types.Matrix, q.Size)
	for i := range matrix {
		matrix[i] = make([]bool, q.Size)
	}

	// place patterns
	placePatterns(q, matrix)

	// place mode and mask data
	placeModeData(matrix, mode, mask)

	// generate FEC data
	// place data and FEC
	writer := newBitReader(matrix, mask, &outOfBounds{QR: q})
	writer.Write(data)

	return matrix
}

func placePatterns(q *QR, matrix types.Matrix) {

	// place corners
	utils.WriteSubmatrix(matrix, types.QRCorner, &types.Point{X: 0, Y: 0})
	utils.WriteSubmatrix(matrix, types.QRCorner, &types.Point{X: q.Size - 7, Y: 0})
	utils.WriteSubmatrix(matrix, types.QRCorner, &types.Point{X: 0, Y: q.Size - 7})

	// place timing pattern
	for c := 8; c < q.Size-7; c += 2 {
		matrix[6][c] = true
		matrix[c][6] = true
	}

	// place valid alignment patterns
	for _, x := range q.AlignmentPatterns {
		for _, y := range q.AlignmentPatterns {
			if !validAlignmentPattern(x, y, q.Size) {
				continue
			}
			utils.WriteSubmatrix(matrix, types.QRCornerSmall, &types.Point{X: x - 2, Y: y - 2})
		}
	}

}

func placeModeData(matrix types.Matrix, mode fec.Mode, mask maskInterface) error {
	toWrite := make([]bool, 15)
	for i, v := range mode.Indicator() {
		toWrite[i] = v
	}
	for i, v := range mask.Indicator() {
		toWrite[i+2] = v
	}
	toWrite, err := utils.XORSlices(toWrite, []bool{
		true, false, true, false, true,
		false, false, false, false, false,
		true, false, false, true, false,
	})
	if err != nil {
		panic(err)
	}

	// upper left
	// lower right
}
