package output

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func MatrixToImage(matrix [][]bool, filepath string) {
	upperLeft, lowerRight := image.Point{0, 0}, image.Point{len(matrix), len(matrix[0])}
	img := image.NewGray(image.Rectangle{upperLeft, lowerRight})

	for x := range matrix {
		for y := range matrix[0] {
			if matrix[x][y] {
				img.Set(x, y, color.Black)
			} else {
				img.Set(x, y, color.White)
			}
		}
	}

	file, err := os.Create(filepath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.Println("Writting result to", filepath)
	png.Encode(file, img)
}

func MatrixToExcel(matrix [][]bool, filepath string) {
	log.Fatal("not imlemented: excel")
}
