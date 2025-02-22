package decoding

func Decode(matrix [][]bool) (string, error) {
	code, err := detectCodeType(matrix)
	if err != nil {
		return "", err
	}
	decoder := newDecoder(matrix, code)
	value, err := decoder.decode()
	if err != nil {
		return "", err
	}

	return value, nil
}
