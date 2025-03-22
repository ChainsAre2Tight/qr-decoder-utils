package cli

import (
	"flag"
	"log"
	"os"

	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/conversion"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/QR/common/masks"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/detection"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/input"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/output"
	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/utils"
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
	outputSizePtr = flag.Int("size", 0, "specifies output matrix size [1-100]")
}

func CLI() {

	// parse master mode argument
	if len(os.Args) < 2 {
		log.Fatal("Master mode not selected [excel | image | mask | decode]")
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

func requireInputName() {
	if *inputFilenamePtr == "" {
		log.Print("Input filename not specified")
		printUsage()
	}
}

// will generate random filename if flag wasn't provided
func requireOutputName() {
	if *outputFilenamePtr == "" {
		*outputFilenamePtr = utils.GenerateRandomFilename()
		log.Printf("WARNING: No --output filename specified, writting to %s", *outputFilenamePtr)
	}
}
func validateOutputSize() {
	if *outputSizePtr < 0 || *outputSizePtr > 100 {
		log.Print("Output size out of range.")
		printUsage()
	}
}

func decode() {
	requireInputName()
	matrix := LoadAndConvert(inputFilenamePtr)

	data, err := decoding.Decode(matrix)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Decoded contents:\n\n", data)
}

func convertImage() {
	requireInputName()
	requireOutputName()
	matrix := LoadAndConvert(inputFilenamePtr)
	output.MatrixToImage(matrix, *outputFilenamePtr)
}

func LoadAndConvert(inputFilenamePtr *string) [][]bool {

	// load image
	img := input.ReadImage(*inputFilenamePtr)

	// detect borders and resize
	validateOutputSize()
	qr, err := detection.DetectCodeOnImage(img, *outputSizePtr)
	if err != nil {
		log.Fatal(err)
	}

	// convert to matrix
	matrix := conversion.ImageToMartix(qr)
	return matrix
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

	matrix := LoadAndConvert(inputFilenamePtr)

	// output to selected format
	outputFunction(matrix, *outputFilenamePtr)
}

func mask() {
	requireOutputName()
	validateOutputSize()

	mask, ok := masks.Masks[*maskPtr]
	if !ok {
		log.Printf("Mask \"%s\" is unknown", *maskPtr)
		printUsage()
	}

	if *outputSizePtr == 0 {
		*outputSizePtr = 21
	}
	result := masks.GenerateMaskedMatrix(*outputSizePtr, mask)

	output.MatrixToExcel(result, *outputFilenamePtr)
}

func printUsage() {
	flag.PrintDefaults()
	os.Exit(1)
}
