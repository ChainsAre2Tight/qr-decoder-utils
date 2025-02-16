package utils

import (
	"reflect"
	"testing"

	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/types"
)

func TestIsSubmatrix(t *testing.T) {
	simpleA := [][]bool{{true, false, true}, {false, true, true}, {true, true, false}}
	simpleB := [][]bool{{true, false, true}, {false, true, true}, {true, true, false}}
	if !IsSubmatrix(simpleA, simpleB, types.NewPoint(0, 0)) {
		t.Fatal("Simple compare fail, shift: 0 0, want: true, outcome: false", simpleA, simpleB)
	}
	if IsSubmatrix(simpleA, simpleB, types.NewPoint(1, 0)) {
		t.Fatal("Simple compare fail, shift 1, 0, want: false, outcome: true", simpleA, simpleB)
	}

	simpleC := [][]bool{
		{true, false, true, false},
		{false, false, false, false},
		{false, false, false, false},
		{false, true, false, true},
	}
	simpleD := [][]bool{
		{true, false},
		{false, false},
	}
	if !IsSubmatrix(simpleC, simpleD, types.NewPoint(0, 0)) {
		t.Fatal("Simple compare fail, shift: 0 0, want: true, outcome: false", simpleC, simpleD)
	}
	if !IsSubmatrix(simpleC, simpleD, types.NewPoint(0, 2)) {
		t.Fatal("Simple compare fail, shift: 0 2, want: true, outcome: false", simpleC, simpleD)
	}
	if IsSubmatrix(simpleC, simpleD, types.NewPoint(1, 0)) {
		t.Fatal("Simple compare fail, shift: 1 0, want: false, outcome: true", simpleC, simpleD)
	}
}

func TestReadMatrixRow(t *testing.T) {
	matrixA := [][]bool{{false, true, false}, {false, false, false}, {true, false, true}, {false, true, false}}

	row2full := []bool{true, false, false, true}
	row1half := []bool{false, true}

	if _, err := ReadMatrixRow(matrixA, 3, 0, 0); err == nil {
		t.Fatal("row 3 is outside of ", matrixA)
	}
	if _, err := ReadMatrixRow(matrixA, 2, 0, 5); err == nil {
		t.Fatal("column 5 is outside of ", matrixA)
	}

	if res, err := ReadMatrixRow(matrixA, 1, 0, 4); err != nil || !reflect.DeepEqual(res, row2full) {
		t.Fatal("want: equal, row: 1, start: 0, end: 4 ", res, err, matrixA, row2full)
	}
	if res, err := ReadMatrixRow(matrixA, 0, 1, 3); err != nil || !reflect.DeepEqual(res, row1half) {
		t.Fatal("want: equal, row: 0, start: 1, end: 3 ", res, err, matrixA, row1half)
	}

}
