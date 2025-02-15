package output

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/utils"
	"github.com/tealeg/xlsx"
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
	name := utils.Concat(filepath, ".xlsx")

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("QR 1")

	if err != nil {
		log.Fatal(err)
	}
	sheet.SetColWidth(0, len(matrix)-1, 2)

	BLACK := xlsx.NewStyle()
	BLACK.Fill = *xlsx.NewFill("solid", "0000000", "00000000")
	BLACK.ApplyFill = true

	for y := range matrix[0] {
		row := sheet.AddRow()

		for x := range matrix {
			cell := row.AddCell()
			if matrix[x][y] {
				cell.SetInt(1)
				cell.SetStyle(BLACK)

			} else {
				cell.SetInt(0)
			}
		}
	}

	log.Println("writting excel to", name)
	err = file.Save(name)
	if err != nil {
		log.Fatal(err)
	}
}
