package input

import (
	"image"
	_ "image/jpeg"
	"log"
	"os"
)

func ReadImage(filepath string) image.Image {
	reader, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()
	log.Println(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y)

	return m
}
