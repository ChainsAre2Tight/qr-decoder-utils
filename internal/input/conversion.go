package input

import (
	"image"
)

func ImageToMartix(image image.Image) [][]bool {
	bounds := image.Bounds()

	matrix := make([][]bool, bounds.Dx())
	for i := range matrix {
		matrix[i] = make([]bool, bounds.Dy())
	}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, _, _, _ := image.At(x, y).RGBA()
			matrix[x-bounds.Min.X][y-bounds.Min.Y] = r < uint32(60000)
		}
	}

	return matrix
}
