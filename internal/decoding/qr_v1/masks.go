package qr_v1

import "github.com/ChainsAre2Tight/qr-decoder-utils/internal/types"

type maskInterface interface {
	At(types.Point) bool
}

type mask000 struct{}
type mask001 struct{}
type mask010 struct{}
type mask011 struct{}
type mask100 struct{}
type mask101 struct{}
type mask110 struct{}
type mask111 struct{}

var masks = map[string]maskInterface{
	"000": mask000{},
	"001": mask001{},
	"010": mask010{},
	"011": mask011{},
	"100": mask100{},
	"101": mask101{},
	"110": mask110{},
	"111": mask111{},
}

func atMatrixXORMask(matrix [][]bool, mask maskInterface, x, y int) bool {
	return matrix[x][y] != mask.At(types.NewPoint(x, y))
}

func (mask000) At(point types.Point) bool {
	return (point.X+point.Y)%2 == 0
}

func (mask001) At(point types.Point) bool {
	return point.Y%2 == 0
}

func (mask010) At(point types.Point) bool {
	return point.X%3 == 0
}

func (mask011) At(point types.Point) bool {
	return (point.X+point.Y)%3 == 0
}

func (mask100) At(point types.Point) bool {
	return (point.Y/2+point.X/3)%2 == 0
}

func (mask101) At(point types.Point) bool {
	return (point.X*point.Y)%2+(point.X*point.Y)%3 == 0
}

func (mask110) At(point types.Point) bool {
	return ((point.X*point.Y)%2 + (point.X*point.Y)%3) == 0
}

func (mask111) At(point types.Point) bool {
	return ((point.X+point.Y)%2+(point.X*point.Y)%3)%2 == 0
}
