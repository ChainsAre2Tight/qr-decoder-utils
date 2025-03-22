package qr

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/interfaces"
)

type blankOOB struct{}

func (blankOOB) SkipCell(x, y int) bool {
	return false
}

func (blankOOB) SkipColumn(x int) bool {
	return false
}

type skipX2Y2 struct{}

func (skipX2Y2) SkipCell(x, y int) bool {
	return y == 2
}

func (skipX2Y2) SkipColumn(x int) bool {
	return x == 2
}

func TestGenReadSequence(t *testing.T) {
	type inputs struct {
		x, y int
		oob  interfaces.OutOfBoundsInterface
	}
	var tests = []struct {
		in  inputs
		out [][2]int
	}{
		{inputs{3, 3, blankOOB{}}, [][2]int{
			{2, 2}, {1, 2}, {2, 1}, {1, 1}, {2, 0}, {1, 0}, {0, 0}, {0, 1}, {0, 2},
		}},
		{inputs{4, 4, blankOOB{}}, [][2]int{
			{3, 3}, {2, 3}, {3, 2}, {2, 2}, {3, 1}, {2, 1}, {3, 0}, {2, 0},
			{1, 0}, {0, 0}, {1, 1}, {0, 1}, {1, 2}, {0, 2}, {1, 3}, {0, 3},
		}},
		{inputs{5, 5, skipX2Y2{}}, [][2]int{
			{4, 4}, {3, 4}, {4, 3}, {3, 3},
			{4, 1}, {3, 1}, {4, 0}, {3, 0},
			{1, 0}, {0, 0}, {1, 1}, {0, 1},
			{1, 3}, {0, 3}, {1, 4}, {0, 4},
		}},
	}
	for _, tt := range tests {
		t.Run(
			fmt.Sprintf("%dx%d, %s", tt.in.x, tt.in.y, reflect.TypeOf(tt.in.oob)),
			func(t *testing.T) {
				result := generateReadSequence(tt.in.x, tt.in.y, tt.in.oob)
				if !reflect.DeepEqual(result, tt.out) {
					t.Error("\ngot ", result, "\nwant", tt.out)
				}
			},
		)
	}
}

type qrv1oob struct{}

func (qrv1oob) SkipCell(x, y int) bool {
	return y == 6 ||
		x <= 8 && y <= 8 ||
		x <= 8 && y >= 13 ||
		x >= 13 && y <= 8
}

func (qrv1oob) SkipColumn(x int) bool {
	return x == 6
}

func TestQRV1Sequence(t *testing.T) {
	sequence := generateReadSequence(21, 21, qrv1oob{})
	expectedSequence := [][2]int{
		// up x3
		{20, 20}, {19, 20}, {20, 19}, {19, 19}, {20, 18}, {19, 18}, {20, 17}, {19, 17},
		{20, 16}, {19, 16}, {20, 15}, {19, 15}, {20, 14}, {19, 14}, {20, 13}, {19, 13},
		{20, 12}, {19, 12}, {20, 11}, {19, 11}, {20, 10}, {19, 10}, {20, 9}, {19, 9},
		// down x3
		{18, 9}, {17, 9}, {18, 10}, {17, 10}, {18, 11}, {17, 11}, {18, 12}, {17, 12},
		{18, 13}, {17, 13}, {18, 14}, {17, 14}, {18, 15}, {17, 15}, {18, 16}, {17, 16},
		{18, 17}, {17, 17}, {18, 18}, {17, 18}, {18, 19}, {17, 19}, {18, 20}, {17, 20},
		// up x3
		{16, 20}, {15, 20}, {16, 19}, {15, 19}, {16, 18}, {15, 18}, {16, 17}, {15, 17},
		{16, 16}, {15, 16}, {16, 15}, {15, 15}, {16, 14}, {15, 14}, {16, 13}, {15, 13},
		{16, 12}, {15, 12}, {16, 11}, {15, 11}, {16, 10}, {15, 10}, {16, 9}, {15, 9},
		// down x3
		{14, 9}, {13, 9}, {14, 10}, {13, 10}, {14, 11}, {13, 11}, {14, 12}, {13, 12},
		{14, 13}, {13, 13}, {14, 14}, {13, 14}, {14, 15}, {13, 15}, {14, 16}, {13, 16},
		{14, 17}, {13, 17}, {14, 18}, {13, 18}, {14, 19}, {13, 19}, {14, 20}, {13, 20},
		// up x3
		{12, 20}, {11, 20}, {12, 19}, {11, 19}, {12, 18}, {11, 18}, {12, 17}, {11, 17},
		{12, 16}, {11, 16}, {12, 15}, {11, 15}, {12, 14}, {11, 14}, {12, 13}, {11, 13},
		{12, 12}, {11, 12}, {12, 11}, {11, 11}, {12, 10}, {11, 10}, {12, 9}, {11, 9},
		// up, up, down, down skipping y = 6
		{12, 8}, {11, 8}, {12, 7}, {11, 7}, {12, 5}, {11, 5}, {12, 4}, {11, 4},
		{12, 3}, {11, 3}, {12, 2}, {11, 2}, {12, 1}, {11, 1}, {12, 0}, {11, 0},
		{10, 0}, {9, 0}, {10, 1}, {9, 1}, {10, 2}, {9, 2}, {10, 3}, {9, 3},
		{10, 4}, {9, 4}, {10, 5}, {9, 5}, {10, 7}, {9, 7}, {10, 8}, {9, 8},
		// down x3
		{10, 9}, {9, 9}, {10, 10}, {9, 10}, {10, 11}, {9, 11}, {10, 12}, {9, 12},
		{10, 13}, {9, 13}, {10, 14}, {9, 14}, {10, 15}, {9, 15}, {10, 16}, {9, 16},
		{10, 17}, {9, 17}, {10, 18}, {9, 18}, {10, 19}, {9, 19}, {10, 20}, {9, 20},
		// last 4, up down up down, skipping x = 6
		{8, 12}, {7, 12}, {8, 11}, {7, 11}, {8, 10}, {7, 10}, {8, 9}, {7, 9},
		{5, 9}, {4, 9}, {5, 10}, {4, 10}, {5, 11}, {4, 11}, {5, 12}, {4, 12},
		{3, 12}, {2, 12}, {3, 11}, {2, 11}, {3, 10}, {2, 10}, {3, 9}, {2, 9},
		{1, 9}, {0, 9}, {1, 10}, {0, 10}, {1, 11}, {0, 11}, {1, 12}, {0, 12},
	}
	if !reflect.DeepEqual(sequence, expectedSequence) {
		t.Error("\ngot\n", sequence, "\nwant\n", expectedSequence)
	}
}

type qrv2oob struct{}

func (qrv2oob) SkipCell(x, y int) bool {
	return y == 6 ||
		x <= 8 && y <= 8 ||
		x <= 8 && y >= 17 ||
		x >= 17 && y <= 8 ||
		x >= 16 && x <= 20 && y >= 16 && y <= 20
}

func (qrv2oob) SkipColumn(x int) bool {
	return x == 6
}

func TestQRV2Sequence(t *testing.T) {

	expectedSequence := [][2]int{
		// up x4
		{24, 24}, {23, 24}, {24, 23}, {23, 23}, {24, 22}, {23, 22}, {24, 21}, {23, 21},
		{24, 20}, {23, 20}, {24, 19}, {23, 19}, {24, 18}, {23, 18}, {24, 17}, {23, 17},
		{24, 16}, {23, 16}, {24, 15}, {23, 15}, {24, 14}, {23, 14}, {24, 13}, {23, 13},
		{24, 12}, {23, 12}, {24, 11}, {23, 11}, {24, 10}, {23, 10}, {24, 9}, {23, 9},
		// down x4
		{22, 9}, {21, 9}, {22, 10}, {21, 10}, {22, 11}, {21, 11}, {22, 12}, {21, 12},
		{22, 13}, {21, 13}, {22, 14}, {21, 14}, {22, 15}, {21, 15}, {22, 16}, {21, 16},
		{22, 17}, {21, 17}, {22, 18}, {21, 18}, {22, 19}, {21, 19}, {22, 20}, {21, 20},
		{22, 21}, {21, 21}, {22, 22}, {21, 22}, {22, 23}, {21, 23}, {22, 24}, {21, 24},
		// up, skip4, up, up
		{20, 24}, {19, 24}, {20, 23}, {19, 23}, {20, 22}, {19, 22}, {20, 21}, {19, 21},
		{20, 15}, {19, 15}, {20, 14}, {19, 14}, {20, 13}, {19, 13}, {20, 12}, {19, 12},
		{20, 11}, {19, 11}, {20, 10}, {19, 10}, {20, 9}, {19, 9},
		// down, down3, skip4, down
		{18, 9}, {17, 9}, {18, 10}, {17, 10}, {18, 11}, {17, 11}, {18, 12}, {17, 12},
		{18, 13}, {17, 13}, {18, 14}, {17, 14}, {18, 15}, {17, 15},
		{18, 21}, {17, 21}, {18, 22}, {17, 22}, {18, 23}, {17, 23}, {18, 24}, {17, 24},
		// up, thin, up minus corner, up...
		{16, 24}, {15, 24}, {16, 23}, {15, 23}, {16, 22}, {15, 22}, {16, 21}, {15, 21},
		{15, 20}, {15, 19}, {15, 18}, {15, 17},
		{15, 16}, {16, 15}, {15, 15}, {16, 14}, {15, 14}, {16, 13}, {15, 13},
	}
	sequence := generateReadSequence(25, 25, qrv2oob{})[:len(expectedSequence)]
	if !reflect.DeepEqual(sequence, expectedSequence) {
		t.Error("\ngot\n", sequence, "\nwant\n", expectedSequence)
	}
}
