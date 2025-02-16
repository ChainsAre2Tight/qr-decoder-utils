package qr_v1_test

import (
	"testing"

	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/qr_v1"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/types"
)

type mockMask struct{}

func (mockMask) At(_ types.Point) bool {
	return false
}

func TestBitReaderSequence(t *testing.T) {
	matrix := make([][]bool, 21)
	for i := range matrix {
		matrix[i] = make([]bool, 21)
	}

	reader := qr_v1.NewBitReader(matrix, mockMask{})

	var expectedSequence = [][2]int{
		{20, 18}, {19, 18}, {20, 17}, {19, 17}, {20, 16}, {19, 16}, {20, 15}, {19, 15},
		{20, 14}, {19, 14}, {20, 13}, {19, 13}, {20, 12}, {19, 12}, {20, 11}, {19, 11},
		{20, 10}, {19, 10}, {20, 9}, {19, 9}, {18, 9}, {17, 9}, {18, 10}, {17, 10},
	}

	for position, pair := range expectedSequence {
		_, x, y := reader.ReadOne()
		if x != pair[0] || y != pair[1] {
			t.Fatalf("Unexpected coordinates at position %d. want: %d, %d, got: %d, %d",
				position, pair[0], pair[1], x, y,
			)
		}
	}
}
