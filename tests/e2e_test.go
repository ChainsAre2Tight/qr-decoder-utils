package e2e_test

import (
	"fmt"
	"testing"

	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/cli"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding"
)

func TestDecoding(t *testing.T) {
	var tests = []struct {
		in   string
		out  string
		size int
	}{
		{"data/qrv1/1.gif", "1", 0},
		{"data/qrv1/time.gif", "time", 0},
		{"data/qrv2/35numbers.gif", "33290056695773232123423233681965212", 0},
		{"data/qrv3/manycharacters.png", "abac4814baba6464cape1551dung4567etch", 0},
		{"data/qrv4/esenin.png", "Жизнь - обман с чарующей тоскою /С. Есенин", 0},
		{"data/qrv2/abaabiaga.gif", "аба абы ага", 0},
		{"data/ECC200/time.png", "time", 12},
		{"data/ECC200/mixer.png", "mixer", 12},
		{"data/ECC200/fallandstand.gif", "Fall seven times and stand up eight", 24},
	}

	for _, tt := range tests {
		t.Run(
			fmt.Sprintf("%s -> %s", tt.in, tt.out),
			func(t *testing.T) {
				matrix := cli.LoadAndConvert(&tt.in, tt.size)
				result, err := decoding.Decode(matrix)
				if err != nil || result != tt.out {
					t.Errorf("\ngot\t%s\nwant\t%s\nerror\t%s", result, tt.out, err)
				}
			},
		)
	}
}
