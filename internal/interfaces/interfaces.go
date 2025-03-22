package interfaces

type OutOfBoundsInterface interface {
	SkipCell(x, y int) bool
	SkipColumn(x int) bool
}

type ModeInterface interface{}

type CodeInterface interface {
	Detect([][]bool) bool
	Decode([][]bool) (string, error)
	Description() string
}

type BitReaderInterface interface {
	ReadMultiple(int) []bool
	ReadBytes() byte
}
