package conversion

import (
	"fmt"
	"image"
	"log"

	"github.com/ChainsAre2Tight/qr-decoder-utils/types"
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

func StripFields(matrix [][]bool) [][]bool {
	// detect upper left
	upperLeft, err := findUpperLeft(matrix)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("upper left is at", upperLeft.X, upperLeft.Y)
	}

	// detect lower left
	lowerLeft := findLowerLeft(matrix)

	// verify left corners are properly aligned and not the same
	if lowerLeft.X != upperLeft.X || lowerLeft.Y == upperLeft.Y {
		log.Fatal("misaligned left border or code is 1 pixel high")
	} else {
		log.Println("lower left is at", lowerLeft.X, lowerLeft.Y)
	}

	// detect upper right
	upperRight := findUpperRight(matrix)

	// verify upper corners are properly aligned and not the same
	if upperLeft.Y != upperRight.Y || upperLeft.X == upperRight.X {
		log.Fatal("misaligned upper border or code is 1 pixel wide")
	} else {
		log.Println("upper right is at", upperRight.X, upperRight.Y)
	}

	// calculate new dimensions
	DX, DY := upperRight.X-upperLeft.X, lowerLeft.Y-upperLeft.Y
	log.Printf("new dimensions are %dx%d", DX, DY)

	// transfer data
	strippedMatrix := make([][]bool, DX)
	for dx := range DX {
		strippedMatrix[dx] = make([]bool, DY)

		x := dx + upperLeft.X

		for dy := range DY {
			y := dy + upperLeft.Y

			strippedMatrix[dx][dy] = matrix[x][y]
		}
	}

	return strippedMatrix
}

func findUpperLeft(matrix [][]bool) (*types.Point, error) {
	for x := range matrix {
		for y := range matrix[0] {
			if matrix[x][y] {
				return &types.Point{X: x, Y: y}, nil
			}
		}
	}
	return &types.Point{X: -1, Y: -1}, fmt.Errorf("image is blank")
}

func findLowerLeft(matrix [][]bool) *types.Point {
	for x := range matrix {
		for y := range matrix[0] {
			if matrix[x][len(matrix[0])-y-1] {
				return &types.Point{X: x, Y: len(matrix[0]) - y - 1}
			}
		}
	}
	return &types.Point{X: -1, Y: -1}
}

func findUpperRight(matrix [][]bool) *types.Point {
	for x := range matrix {
		for y := range matrix[0] {
			if matrix[len(matrix)-x-1][y] {
				return &types.Point{X: len(matrix) - x - 1, Y: y}
			}
		}
	}
	return &types.Point{X: -1, Y: -1}
}
