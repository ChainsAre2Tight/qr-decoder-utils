package fec

type Mode interface {
	GenerateErrorCorrection(data []bool) ([]bool, error)
	Indicator() []bool
}

type modeL struct{}
type modeM struct{}
type modeQ struct{}
type modeH struct{}

var Modes = map[string]Mode{
	"L": &modeL{},
	"M": &modeM{},
	"Q": &modeQ{},
	"H": &modeH{},
}

func (m modeL) GenerateErrorCorrection(data []bool) ([]bool, error) {
	panic("unimplemented")
}

func (m modeM) GenerateErrorCorrection(data []bool) ([]bool, error) {
	panic("unimplemented")
}

func (m modeQ) GenerateErrorCorrection(data []bool) ([]bool, error) {
	panic("unimplemented")
}

func (m modeH) GenerateErrorCorrection(data []bool) ([]bool, error) {
	panic("unimplemented")
}

func (m modeL) Indicator() []bool {
	panic("unimplemented")
}

func (m modeM) Indicator() []bool {
	panic("unimplemented")
}

func (m modeQ) Indicator() []bool {
	panic("unimplemented")
}

func (m modeH) Indicator() []bool {
	panic("unimplemented")
}
