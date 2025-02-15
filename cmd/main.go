package main

import (
	"flag"
	"log"

	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/conversion"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/detection"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/input"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/output"
)

func main() {
	// parse arguments
	inputFilenamePtr := flag.String("input", "./image.png", "specifies an image file to parse")
	outputFilenamePtr := flag.String("output", "./result", "specifies an output file name")
	outputModePtr := flag.String("mode", "excel", "specifies output mode, can be excel or image")
	flag.Parse()

	log.Println("Reading from", *inputFilenamePtr, "and writting to", *outputFilenamePtr)

	var outputFunction func([][]bool, string)
	switch *outputModePtr {
	case "excel":
		log.Println("output is set as an excel spreadsheet")
		outputFunction = output.MatrixToExcel
	case "image":
		log.Println("output is set as an image")
		outputFunction = output.MatrixToImage
	default:
		log.Fatal("unknown mode: ", *outputModePtr)
	}

	// load image
	img := input.ReadImage(*inputFilenamePtr)

	// detect qr and resize
	qr, err := detection.DetectQR(img)
	if err != nil {
		log.Fatal(err)
	}

	// convert to matrix
	matrix := conversion.ImageToMartix(qr)

	// output to selected format
	outputFunction(matrix, *outputFilenamePtr)
}
