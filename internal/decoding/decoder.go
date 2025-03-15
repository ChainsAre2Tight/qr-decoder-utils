package decoding

import (
	"log"
	"reflect"

	bitreader "github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/common/bit_reader"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/interfaces"
)

type decoder struct {
	matrix [][]bool
	code   interfaces.CodeInterface
}

func newDecoder(matrix [][]bool, code interfaces.CodeInterface) decoder {
	return decoder{
		code:   code,
		matrix: matrix,
	}
}

func (d *decoder) decode() (string, error) {
	_, mask, err := d.code.ReadMetadata(d.matrix)
	if err != nil {
		return "", err
	}

	log.Println("Detected mask:", reflect.TypeOf(mask).Name())
	reader := bitreader.NewBitReader(d.matrix, mask, d.code.OOB())

	format, err := d.code.ReadFormat(d.matrix, mask, reader)
	if err != nil {
		return "", err
	}
	log.Println("Detected format:", reflect.TypeOf(format).Name())

	data, err := format.ReadData(d.matrix, mask, reader)
	if err != nil {
		return "", err
	}
	return data, nil
}
