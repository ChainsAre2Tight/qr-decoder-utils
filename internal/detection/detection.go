package detection

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math"

	"golang.org/x/image/draw"
)

func DetectQR(img image.Image, sizeOverride int) (image.Image, error) {
	border, err := detectBorders(img)
	if err != nil {
		return image.Black, err
	}
	cropped, err := cropFields(img, border)
	if err != nil {
		return image.Black, err
	}
	var newX, newY int
	if sizeOverride == 0 {
		log.Println("Size not specified, attempting to detect...")
		log.Println("WARNING: current implementation allows only for QR code size detecion, other code formats are not supported")

		pixelSize, err := detectPixelSize(cropped)
		if err != nil {
			return image.Black, err
		}

		log.Println("Pixel size is", pixelSize)

		newX, newY = calculateNewDimensions(cropped, pixelSize)
		log.Printf("Calculated dimesions: %dx%d", newX, newY)
		log.Println("INFO: if determined code dimensions are wrong, force them with --size")
	} else {
		log.Println("Provided size override")
		newX = sizeOverride
		newY = sizeOverride
	}

	log.Printf("Converting to QR %dx%d ", newX, newY)

	resized := resize(cropped, newX, newY)
	return resized, nil
}

func detectBorders(img image.Image) (image.Rectangle, error) {
	// detect upper left
	ul, err := detectUpperLeft(img)
	if err != nil {
		return image.Rectangle{}, err
	}
	// log.Println("ul is at", ul.X, ul.Y)

	ll, err := detectLowerLeft(img, ul)
	if err != nil {
		return image.Rectangle{}, err
	}
	// log.Println("bl is at", ll.X, ll.Y)

	ur, err := detectUpperRight(img, ul)
	if err != nil {
		return image.Rectangle{}, err
	}
	// log.Println("ur is at", ur.X, ur.Y)

	lr, err := detectLowerRightFromLowerLeft(img, ll)
	if err != nil {
		return image.Rectangle{}, err
	}
	// log.Println("lr is at", lr.X, lr.Y)

	borders := calcBorder(ul, ur, ll, lr)
	log.Println("Caclulated borders:", borders)
	return borders, nil
}

func cropFields(img image.Image, border image.Rectangle) (image.Image, error) {
	type subImager interface {
		SubImage(r image.Rectangle) image.Image
	}

	simg, ok := img.(subImager)
	if !ok {
		return nil, fmt.Errorf("image does not support cropping")
	}

	return simg.SubImage(border), nil
}

// Detects pixel size of a QR code by scanning diagonally from upper-right corner
// of the code until it comes across an all-white inverted L shape
func detectPixelSize(img image.Image) (int, error) {
	bounds := img.Bounds()

	for d := 1; d < bounds.Dx()-5 && d < bounds.Dy()-5; d++ {
		x, y := bounds.Max.X-d, bounds.Min.Y+d

		value := rgbaToValue(img.At(x, y))
		for f := range 5 {
			value += rgbaToValue(img.At(x-f-1, y))
			value += rgbaToValue(img.At(x, y+f+1))
		}

		if value > 500000 {
			return d, nil
		}
	}
	return -1, fmt.Errorf("no pixel size detected")
}

func calculateNewDimensions(img image.Image, pixelSize int) (int, int) {
	bounds := img.Bounds()
	resultX := calcOneDimension(bounds.Dx(), pixelSize)
	resultY := calcOneDimension(bounds.Dy(), pixelSize)
	return resultX, resultY
}

func calcOneDimension(raw, pixelSize int) int {
	return int(math.Ceil(float64(raw) / float64(pixelSize)))
}

func resize(img image.Image, newX, newY int) image.Image {
	result := image.NewRGBA(image.Rect(0, 0, newX, newY))
	draw.NearestNeighbor.Scale(result, result.Rect, img, img.Bounds(), draw.Over, nil)
	return result
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

func detectLowerRightFromLowerLeft(img image.Image, lowerLeft image.Point) (image.Point, error) {
	bounds := img.Bounds()

	for x := bounds.Max.X; x > lowerLeft.X+5; x-- {
		value := rgbaToValue(img.At(x, lowerLeft.Y))
		for d := range 5 {
			value += rgbaToValue(img.At(x, lowerLeft.Y-d-1))
			value += rgbaToValue(img.At(x-d-1, lowerLeft.Y))
		}
		if isCorner(value) {
			return image.Point{x, lowerLeft.Y}, nil
		}
	}
	return image.Point{0, 0}, fmt.Errorf("no corner detected")
}

func calcBorder(ul, ur, ll, lr image.Point) image.Rectangle {
	return image.Rectangle{
		Min: ul,
		Max: image.Point{
			X: max(ur.X, lr.X),
			Y: ll.Y,
		},
	}
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
