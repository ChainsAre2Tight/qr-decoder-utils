package types

type Point struct {
	X, Y int
}

func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

var QRCorner = [][]bool{
	{true, true, true, true, true, true, true},
	{true, false, false, false, false, false, true},
	{true, false, true, true, true, false, true},
	{true, false, true, true, true, false, true},
	{true, false, true, true, true, false, true},
	{true, false, false, false, false, false, true},
	{true, true, true, true, true, true, true},
}
