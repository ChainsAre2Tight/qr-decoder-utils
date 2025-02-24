package main

import (
	"flag"
	"log"
	"os"

	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/conversion"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/common/masks"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/detection"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/input"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/output"
)

var (
	inputFilenamePtr  *string
	outputFilenamePtr *string
	includeMasksPtr   *bool
	outputSizePtr     *int
	maskPtr           *string
)

func init() {
	inputFilenamePtr = flag.String("input", "", "specifies an image file to parse")
	outputFilenamePtr = flag.String("output", "", "specifies an output file name")
	includeMasksPtr = flag.Bool("include-masks", false, "include all known masks as additional sheets")
	maskPtr = flag.String("mask", "None", "specifies mask. [000-111]")
	outputSizePtr = flag.Int("size", 21, "specifies output matrix size [1-100]")
}

func main() {

	// parse master mode argument
	if len(os.Args) < 2 {
		log.Fatal("Master mode not selected \n--mode [excel | image | mask | decode]")
	}
	masterModePtr := os.Args[1]

	// removes first argument so that flag.Parse doesn't get stuck at first positional arg
	os.Args = os.Args[1:]
	flag.Parse()

	switch masterModePtr {
	case "excel":
		convertExcel()
	case "image":
		convertImage()
	case "mask":
		mask()
	case "decode":
		decode()
	default:
		log.Printf("Unknown master mode: %s", masterModePtr)
		printUsage()
	}

}

func decode() {

	if *inputFilenamePtr == "" {
		log.Print("Input filename not specified")
		printUsage()
	}

	log.Println("Reading from", *inputFilenamePtr, "and attempting to decode")

	// load image
	img := input.ReadImage(*inputFilenamePtr)

	// detect borders and resize
	qr, err := detection.DetectQR(img)
	if err != nil {
		log.Fatal(err)
	}

	// convert to matrix
	matrix := conversion.ImageToMartix(qr)

	data, err := decoding.Decode(matrix)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Decoded contents:\n\n", data)
}

func convertImage() {
	log.Print("output is set as to image")
	matrix := loadAndConvert(inputFilenamePtr, outputFilenamePtr)
	output.MatrixToImage(matrix, *outputFilenamePtr)
}

func loadAndConvert(inputFilenamePtr, outputFilenamePtr *string) [][]bool {
	if *inputFilenamePtr == "" {
		log.Print("Input filename not specified")
		printUsage()
	}
	if *outputFilenamePtr == "" {
		log.Print("Output filename not specified")
		printUsage()
	}
	log.Println("Reading from", *inputFilenamePtr, "and writting to", *outputFilenamePtr)

	// load image
	img := input.ReadImage(*inputFilenamePtr)

	// detect borders and resize
	qr, err := detection.DetectQR(img)
	if err != nil {
		log.Fatal(err)
	}

	// convert to matrix
	matrix := conversion.ImageToMartix(qr)
	return matrix
}

func convertExcel() {
	flag.Parse()

	log.Print("output is set as to excel spreadsheet")
	var outputFunction func([][]bool, string)
	if *includeMasksPtr {
		outputFunction = output.MatrixToExcelWithMasks
	} else {
		outputFunction = output.MatrixToExcel
	}

	matrix := loadAndConvert(inputFilenamePtr, outputFilenamePtr)

	// output to selected format
	outputFunction(matrix, *outputFilenamePtr)
}

func mask() {
	if *outputSizePtr < 1 || *outputSizePtr > 100 {
		log.Print("Output size out of range.")
		printUsage()
	}

	if *outputFilenamePtr == "" {
		log.Print("Output filename not specified")
		printUsage()
	}

	mask, ok := masks.Masks[*maskPtr]
	if !ok {
		log.Printf("Mask \"%s\" is unknown", *maskPtr)
		printUsage()
	}

	result := masks.GenerateMaskedMatrix(*outputSizePtr, mask)

	output.MatrixToExcel(result, *outputFilenamePtr)
}

func printUsage() {
	flag.PrintDefaults()
	os.Exit(1)
}
