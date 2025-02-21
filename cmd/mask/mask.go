package main

import (
	"flag"
	"log"

	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/common/masks"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/output"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/types"
)

func main() {
	outputMaskPtr := flag.String("mask", "000", "specifies mask range 000-111")
	outputFilenamePtr := flag.String("output", "./result", "specifies an output file name")
	outputSizePtr := flag.Int("size", 21, "specifies output matrix size")

	flag.Parse()
	// log.Printf("Will output mask %d of size %f into %d", *outputMaskPtr, *outputSizePtr, *outputFilenamePtr)
	mask, ok := masks.Masks[*outputMaskPtr]
	if !ok {
		log.Fatal("Mask is unknown")
	}

	matrix := generate_matrix(*outputSizePtr, mask)

	output.MatrixToExcel(matrix, *outputFilenamePtr)
}

func generate_matrix(size int, mask masks.MaskInterface) [][]bool {
	result := make([][]bool, size)
	for i := range size {
		result[i] = make([]bool, size)
		for j := range size {
			result[i][j] = mask.At(types.NewPoint(i, j))
		}
	}
	return result
}
