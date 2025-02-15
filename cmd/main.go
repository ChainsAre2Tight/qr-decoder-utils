package main

import (
	"flag"
	"log"

	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/detection"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/input"
)

func main() {
	// parse arguments
	inputFilenamePtr := flag.String("input", "./image.png", "specifies an image file to parse")
	outputFilenamePtr := flag.String("output", "./result", "specifies an output file name")
	outputModePtr := flag.String("mode", "excel", "specifies output mode, can be excel or image")
	flag.Parse()

	log.Println("Reading from", *inputFilenamePtr, "and writting to", *outputFilenamePtr)

	// var outputFunction func([][]bool, *string)
	switch *outputModePtr {
	case "excel":
		log.Println("output is set as an excel spreadsheet")
		// outputFunction = output.MatrixToExcel
	case "image":
		log.Println("output is set as an image")
		// outputFunction = output.MatrixToImage
	default:
		log.Fatal("unknown mode: ", *outputModePtr)
	}

	// load image
	img := input.ReadImage(*inputFilenamePtr)
	border, err := detection.DetectQR(img)
	if err != nil {
		log.Fatal(err)
	}
	cropped, err := detection.CropFields(img, border)
	if err != nil {
		log.Fatal(err)
	}
	pixelSize, err := detection.DetectPixelSize(cropped)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("pizel size is ", pixelSize)
	// detection.Wrtout(cropped, image.Point{pixelSize, pixelSize})

	newDimensions := detection.CalculateNewDimensions(cropped, pixelSize)
	log.Println("Converting to QR of size ", newDimensions)

	resized := detection.Resize(cropped, newDimensions)
	detection.Wrtout(resized)

	// // convert to matrix
	// matrix := conversion.ImageToMartix(image)
	// // strip fields
	// strippedMatrix := conversion.StripFields(matrix)
	// outputFunction(strippedMatrix, outputFilenamePtr)
	// // resize
	// _ = conversion.ResizeMatrix(strippedMatrix, 21)
	// // output to .xlsx
}
