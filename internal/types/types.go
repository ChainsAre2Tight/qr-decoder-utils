package types

type Point struct {
	X, Y int
}

func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

// Character Count Indicator length
// Refer to table 3
type CCI struct {
	Numeric      int
	Alphanumeric int
	Byte         int
	Kanji        int
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

var QRCornerSmall = [][]bool{
	{true, true, true, true, true},
	{true, false, false, false, true},
	{true, false, true, false, true},
	{true, false, false, false, true},
	{true, true, true, true, true},
}
