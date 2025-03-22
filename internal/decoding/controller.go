package decoding

import (
	"log"
)

func Decode(matrix [][]bool) (string, error) {
	code, err := detectCodeType(matrix)
	if err != nil {
		return "", err
	}
	log.Println("Detected code type:", code.Description())
	value, err := code.Decode(matrix)
	if err != nil {
		return "", err
	}

	return value, nil
}
