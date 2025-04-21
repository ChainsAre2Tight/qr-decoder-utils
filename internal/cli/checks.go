package cli

import (
	"log"

	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/utils"
)

func requireInputName() {
	if *inputFilenamePtr == "" {
		log.Print("Input not specified")
		printUsage()
	}
}

func requireCodeType() {
	if *codePtr == "" {
		log.Println("Code type not specififed")
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
