package main

import (
	"flag"
	"log"
	"os"
	"reflect"

	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/conversion"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/common/masks"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/detection"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/input"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/output"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/types"
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

	// detect code type
	code, err := decoding.DetectCodeType(matrix)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Detected code type:", reflect.TypeOf(code).Name())

	// attempt to decode
	data, err := code.Decode(matrix)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("the data is: ", data)
}

func convert() {
	modePtr := flag.String("format", "", "specifies conversion mode [excel | image]")
	inputFilenamePtr := flag.String("input", "", "specifies an image file to parse")
	outputFilenamePtr := flag.String("output", "", "specifies an output file name")
	flag.Parse()

	var outputFunction func([][]bool, string)
	switch *modePtr {
	case "excel":
		log.Print("output is set as to excel spreadsheet")
		outputFunction = output.MatrixToExcel
	case "image":
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

	result := make([][]bool, *outputSizePtr)
	for i := range *outputSizePtr {
		result[i] = make([]bool, *outputSizePtr)
		for j := range *outputSizePtr {
			result[i][j] = mask.At(types.NewPoint(i, j))
		}
	}

	output.MatrixToExcel(result, *outputFilenamePtr)
}

func printUsage() {
	flag.PrintDefaults()
	os.Exit(1)
}
