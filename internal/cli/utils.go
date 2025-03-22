package cli

import (
	"flag"
	"log"
	"os"

	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/input"
)

func LoadAndConvert(inputFilenamePtr *string) [][]bool {

	// load image
	img := input.ReadImage(*inputFilenamePtr)

	// detect borders and resize
	validateOutputSize()
	qr, err := input.DetectCodeOnImage(img, *outputSizePtr)
	if err != nil {
		log.Fatal(err)
	}

	// convert to matrix
	matrix := input.ImageToMartix(qr)
	return matrix
}

func printUsage() {
	flag.PrintDefaults()
	os.Exit(1)
}
