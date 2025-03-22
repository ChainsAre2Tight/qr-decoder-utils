package bitreader

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
				result := GenerateReadSequence(tt.in.x, tt.in.y, tt.in.oob)
				if !reflect.DeepEqual(result, tt.out) {
					t.Error("\ngot ", result, "\nwant", tt.out)
				}
			},
		)
	}
}
