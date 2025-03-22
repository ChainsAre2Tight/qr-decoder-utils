package qr_v1

type oob struct{}

func (oob) SkipCell(x, y int) bool {
	return y == 6 ||
		x <= 8 && y <= 8 ||
		x <= 8 && y >= 13 ||
		x >= 13 && y <= 8
}

func (oob) SkipColumn(x int) bool {
	return x == 6
}
