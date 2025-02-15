package utils

import (
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
