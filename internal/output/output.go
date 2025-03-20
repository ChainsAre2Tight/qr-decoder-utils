package output

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"reflect"

	"github.com/ChainsAre2Tight/qr-decoder-utils/internal/decoding/QR/common/masks"
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

	name := utils.ForceFileExtension(filepath, ".png")
	file, err := os.Create(name)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.Println("Writting result to", name)
	png.Encode(file, img)
}

func MatrixToExcel(matrix [][]bool, filepath string) {
	name := utils.ForceFileExtension(filepath, ".xlsx")

	file := xlsx.NewFile()
	mainSheet, err := file.AddSheet("QR 1")
	if err != nil {
		log.Fatal(err)
	}

	matrixToSheet(matrix, mainSheet)

	log.Println("Writing code to", name)
	err = file.Save(name)
	if err != nil {
		log.Fatal(err)
	}
}

func MatrixToExcelWithMasks(matrix [][]bool, filepath string) {
	name := utils.Concat(filepath, ".xlsx")

	file := xlsx.NewFile()
	mainSheet, err := file.AddSheet("QR")
	if err != nil {
		log.Fatal(err)
	}
	matrixToSheet(matrix, mainSheet)

	for _, mask := range masks.Masks {
		maskSheet, err := file.AddSheet(reflect.TypeOf(mask).Name())
		if err != nil {
			log.Fatal(err)
		}
		matrixToSheet(masks.GenerateMaskedMatrix(len(matrix), mask), maskSheet)
	}

	log.Println("Writing code and masks to", name)
	err = file.Save(name)
	if err != nil {
		log.Fatal(err)
	}
}

func matrixToSheet(matrix [][]bool, sheet *xlsx.Sheet) {
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
}
