package main

import (
	"flag"
	"log"

	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/conversion"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/input"
)

func main() {
	// parse arguments
	inputFilenamePtr := flag.String("input", "image.png", "specifies an image file to parse")
	outputFilenamePtr := flag.String("output", "result.xlsx", "specifies an output file name")
	flag.Parse()

	log.Println("Reading from", *inputFilenamePtr, "and writting to", *outputFilenamePtr)

	// load image
	image := input.ReadImage(*inputFilenamePtr)

	// convert to matrix
	matrix := conversion.ImageToMartix(image)
	// strip fields
	_ = conversion.StripFields(matrix)
	// resize
	// output to .xlsx
}
