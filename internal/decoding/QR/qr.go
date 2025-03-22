package qr

// Character Count Indicator length
type cci struct {
	Numeric      int
	Alphanumeric int
	Byte         int
	Kanji        int
}

// refer to table 3
var CCI1dash9 = &cci{Numeric: 10, Alphanumeric: 9, Byte: 8, Kanji: 8}
var CCI10dash26 = &cci{Numeric: 12, Alphanumeric: 11, Byte: 16, Kanji: 10}
var CCI27dash40 = &cci{Numeric: 14, Alphanumeric: 13, Byte: 16, Kanji: 12}

type QR struct {
	Name              string
	Size              int
	Cci               *cci
	AlignmentPatterns []int
}

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
