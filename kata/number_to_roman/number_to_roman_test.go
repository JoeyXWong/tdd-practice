package number_to_roman

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func convertToRoman(input int) string {
	return NewNumberToRoman().Convert(input)
}

func TestNumberToRoman_Convert(t *testing.T) {
	Convey("Test: Number should be converted to roman numeral", t, func() {
		So(convertToRoman(0), ShouldEqual, "")
		So(convertToRoman(1), ShouldEqual, "I")
		So(convertToRoman(2), ShouldEqual, "II")
		So(convertToRoman(3), ShouldEqual, "III")
		So(convertToRoman(5), ShouldEqual, "V")
		So(convertToRoman(6), ShouldEqual, "VI")
		So(convertToRoman(10), ShouldEqual, "X")
		So(convertToRoman(15), ShouldEqual, "XV")
		So(convertToRoman(15), ShouldEqual, "")

	})
}

type numberToRoman struct {
}

type NumberToRoman interface {
	Convert(num int) string
}

func NewNumberToRoman() NumberToRoman {
	return &numberToRoman{}
}

func (numberToRoman) Convert(num int) string {
	roman := ""
	for num > 0 {
		roman += GetRomanNumeral(&num)
	}
	return roman
}

func GetRomanNumeral(num *int) string {
	switch {
	case *num >= 10:
		*num -= 10
		return "X"
	case *num >= 5:
		*num -= 5
		return "V"
	case *num >= 1:
		*num -= 1
		return "I"
	}
	return ""
}
