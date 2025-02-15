package detection

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"

	"golang.org/x/image/draw"
)

func DetectQR(img image.Image) (image.Rectangle, error) {
	// detect upper left
	ul, err := detectUpperLeft(img)
	if err != nil {
		return image.Rectangle{}, err
	}
	log.Println("ul is at", ul.X, ul.Y)
	ll, err := detectLowerLeft(img, ul)
	if err != nil {
		return image.Rectangle{}, err
	}
	log.Println("bl is at", ll.X, ll.Y)

	ur, err := detectUpperRight(img, ul)
	if err != nil {
		return image.Rectangle{}, err
	}
	log.Println("ur is at", ur.X, ur.Y)

	return image.Rectangle{Min: ul, Max: image.Point{X: ur.X, Y: ll.Y}}, nil
}

func CropFields(img image.Image, border image.Rectangle) (image.Image, error) {
	type subImager interface {
		SubImage(r image.Rectangle) image.Image
	}

	simg, ok := img.(subImager)
	if !ok {
		return nil, fmt.Errorf("image does not support cropping")
	}

	return simg.SubImage(border), nil
}

func DetectPixelSize(img image.Image) (int, error) {
	bounds := img.Bounds()

	for d := 0; d < bounds.Max.X-5 && d < bounds.Max.Y-5; d++ {
		x, y := bounds.Min.X+d, bounds.Min.Y+d

		value := rgbaToValue(img.At(x, y))
		for d := range 5 {
			value += rgbaToValue(img.At(x+d+1, y))
			value += rgbaToValue(img.At(x, y+d+1))
		}

		if value > 500000 {
			return d, nil
		}
	}
	return -1, fmt.Errorf("no pixel size detected")
}

func CalculateNewDimensions(img image.Image, pixelSize int) int {
	bounds := img.Bounds()
	result := int(math.Ceil(float64(bounds.Dx()) / float64(pixelSize)))
	return result
}

func Resize(img image.Image, newDimensions int) image.Image {
	result := image.NewRGBA(image.Rect(0, 0, newDimensions, newDimensions))
	draw.NearestNeighbor.Scale(result, result.Rect, img, img.Bounds(), draw.Over, nil)
	return result
}

func Wrtout(img image.Image) {
	bounds := img.Bounds()
	img2 := image.NewNRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {

			img2.Set(x, y, img.At(x, y))

		}
	}

	file, err := os.Create("./data/result2.png")
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(file, img2)
}

func isCorner(value int) bool {
	return value < 300000
}

func detectUpperLeft(img image.Image) (image.Point, error) {
	bounds := img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y-5; y++ {
		for x := bounds.Min.X; x < bounds.Max.X-5; x++ {
			value := rgbaToValue(img.At(x, y))

			for d := range 5 {
				value += rgbaToValue(img.At(x+d+1, y))
				value += rgbaToValue(img.At(x, y+d+1))
			}

			if isCorner(value) {
				return image.Point{x, y}, nil
			}
		}
	}
	return image.Point{0, 0}, fmt.Errorf("no corner detected")
}

func detectLowerLeft(img image.Image, upperLeft image.Point) (image.Point, error) {
	bounds := img.Bounds()

	for y := bounds.Max.Y; y > upperLeft.Y+5; y-- {
		value := rgbaToValue(img.At(upperLeft.X, y))
		for d := range 5 {
			value += rgbaToValue(img.At(upperLeft.X, y-d-1))
			value += rgbaToValue(img.At(upperLeft.X+d+1, y))
		}
		if isCorner(value) {
			return image.Point{upperLeft.X, y}, nil
		}
	}
	return image.Point{0, 0}, fmt.Errorf("no corner detected")
}

func detectUpperRight(img image.Image, upperLeft image.Point) (image.Point, error) {
	bounds := img.Bounds()

	for x := bounds.Max.X; x > upperLeft.X+5; x-- {
		value := rgbaToValue(img.At(x, upperLeft.Y))
		for d := range 5 {
			value += rgbaToValue(img.At(x, upperLeft.Y+d+1))
			value += rgbaToValue(img.At(x-d-1, upperLeft.Y))
		}
		if isCorner(value) {
			return image.Point{x, upperLeft.Y}, nil
		}
	}
	return image.Point{0, 0}, fmt.Errorf("no corner detected")
}

func rgbaToValue(pixel color.Color) int {
	r, g, b, _ := pixel.RGBA()
	return int((r + g + b) / 3)
}
