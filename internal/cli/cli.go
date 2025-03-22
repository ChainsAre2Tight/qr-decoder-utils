package cli

import (
	"flag"
	"log"
	"os"
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
