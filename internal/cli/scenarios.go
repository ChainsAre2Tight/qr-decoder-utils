package cli

import (
	"log"

	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding"
	qr "github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/QR"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/output"
)

func convertImage() {
	requireInputName()
	requireOutputName()
	matrix := LoadAndConvert(inputFilenamePtr, *outputSizePtr)
	output.MatrixToImage(matrix, *outputFilenamePtr)
}

func decode() {
	requireInputName()
	matrix := LoadAndConvert(inputFilenamePtr, *outputSizePtr)

	data, err := decoding.Decode(matrix)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Decoded contents:\n\n", data)
}

func convertExcel() {
	requireInputName()
	requireOutputName()

	var outputFunction func([][]bool, string)
	if *includeMasksPtr {
		outputFunction = output.MatrixToExcelWithMasks
	} else {
		outputFunction = output.MatrixToExcel
	}

	matrix := LoadAndConvert(inputFilenamePtr, *outputSizePtr)

	// output to selected format
	outputFunction(matrix, *outputFilenamePtr)
}

func mask() {
	requireOutputName()
	validateOutputSize()

	mask, ok := qr.Masks[*maskPtr]
	if !ok {
		log.Printf("Mask \"%s\" is unknown", *maskPtr)
		printUsage()
	}

	if *outputSizePtr == 0 {
		*outputSizePtr = 21
	}
	result := qr.GenerateMaskedMatrix(*outputSizePtr, mask)

	output.MatrixToExcel(result, *outputFilenamePtr)
}
