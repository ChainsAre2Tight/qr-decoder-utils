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

func main() {

	// parse master mode argument
	if len(os.Args) < 2 {
		log.Fatal("Master mode not selected \n--mode [convert | mask | decode]")
	}
	masterModePtr := os.Args[1]

	// removes first argument so that flag.Parse doesn't get stuck at first positional arg
	os.Args = os.Args[1:]

	switch masterModePtr {
	case "convert":
		convert()
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
	inputFilenamePtr := flag.String("input", "", "specifies an image file to parse")
	flag.Parse()

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
	log.Println("the data is: ", data)
}

func convert() {
	modePtr := flag.String("format", "", "specifies conversion mode [excel | image]")
	inputFilenamePtr := flag.String("input", "", "specifies an image file to parse")
	outputFilenamePtr := flag.String("output", "", "specifies an output file name")
	includeMasksPtr := flag.Bool("include-masks", false, "include all known masks as additional sheets")
	flag.Parse()

	var outputFunction func([][]bool, string)
	switch *modePtr {
	case "excel":
		log.Print("output is set as to excel spreadsheet")
		if *includeMasksPtr {
			outputFunction = output.MatrixToExcelWithMasks
		} else {
			outputFunction = output.MatrixToExcel
		}
	case "image":
		if *includeMasksPtr {
			log.Print("--include-masks module is supported only for excel conversion")
			printUsage()
		}
		log.Print("output is set as to image")
		outputFunction = output.MatrixToImage
	default:
		log.Printf("unknown format: \"%s\"", *modePtr)
		printUsage()
	}

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

	// output to selected format
	outputFunction(matrix, *outputFilenamePtr)
}

func mask() {
	maskPtr := flag.String("mask", "None", "specifies mask. [000-111]")
	outputFilenamePtr := flag.String("output", "", "specifies an output file name")
	outputSizePtr := flag.Int("size", 21, "specifies output matrix size [1-100]")
	flag.Parse()

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
