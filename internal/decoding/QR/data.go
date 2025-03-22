package qr

// All hardcoded qr codes used in this app
// Refer to table 1 for size parameter,
// Refer to table 3 for content length parameters and
// table E.1 for list of alignment patterns
var QR_CODES = []*QR{
	{Name: "QR Version 1", Size: 21, Cci: CCI1dash9, AlignmentPatterns: []int{}},
	{Name: "QR Version 2", Size: 25, Cci: CCI1dash9, AlignmentPatterns: []int{6, 18}},
	{Name: "QR Version 3", Size: 29, Cci: CCI1dash9, AlignmentPatterns: []int{6, 22}},
	{Name: "QR Version 4", Size: 33, Cci: CCI1dash9, AlignmentPatterns: []int{6, 26}},
}

func DetectQR(matrix [][]bool) (*QR, bool) {
	for _, code := range QR_CODES {
		ok := code.Detect(matrix)
		if ok {
			return code, true
		}
	}
	return nil, false
}
