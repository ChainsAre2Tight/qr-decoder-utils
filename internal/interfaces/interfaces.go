package interfaces

type ModeInterface interface{}

type CodeInterface interface {
	Detect([][]bool) bool
	Decode([][]bool) (string, error)
	Description() string
}
