package decoding

import (
	"log"
	"reflect"
)

func Decode(matrix [][]bool) (string, error) {
	code, err := detectCodeType(matrix)
	if err != nil {
		return "", err
	}
	log.Println("Detected code type:", reflect.TypeOf(code).Name())
	decoder := newDecoder(matrix, code)
	value, err := decoder.decode()
	if err != nil {
		return "", err
	}

	return value, nil
}
