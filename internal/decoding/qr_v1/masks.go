package qr_v1

import "github.com/ChainsAre2Tight/qr-decoder-utils/internal/types"

type maskInterface interface {
	At(types.Point) bool
}

type Mask000 struct{}
type Mask001 struct{}
type Mask010 struct{}
type Mask011 struct{}
type Mask100 struct{}
type Mask101 struct{}
type Mask110 struct{}
type Mask111 struct{}

var masks = map[string]maskInterface{
	"000": Mask000{},
	"001": Mask001{},
	"010": Mask010{},
	"011": Mask011{},
	"100": Mask100{},
	"101": Mask101{},
	"110": Mask110{},
	"111": Mask111{},
}

func atMatrixXORMask(matrix [][]bool, mask maskInterface, x, y int) bool {
	return matrix[x][y] != mask.At(types.NewPoint(x, y))
}

func (Mask000) At(point types.Point) bool {
	return (point.X+point.Y)%2 == 0
}

func (Mask001) At(point types.Point) bool {
	return point.Y%2 == 0
}

func (Mask010) At(point types.Point) bool {
	return point.X%3 == 0
}

func (Mask011) At(point types.Point) bool {
	return (point.X+point.Y)%3 == 0
}

func (Mask100) At(point types.Point) bool {
	return (point.Y/2+point.X/3)%2 == 0
}

func (Mask101) At(point types.Point) bool {
	return (point.X*point.Y)%2+(point.X*point.Y)%3 == 0
}

func (Mask110) At(point types.Point) bool {
	return ((point.X*point.Y)%2 + (point.X*point.Y)%3) == 0
}

func (Mask111) At(point types.Point) bool {
	return ((point.X+point.Y)%2+(point.X*point.Y)%3)%2 == 0
}
