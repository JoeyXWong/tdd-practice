package string_to_atoi

import (
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"strconv"
	"strings"
	"testing"
	"unicode"
)

func getConverted(input string) int {
	return NewStringToAtoi().Convert(input)
}

func TestStringToAtoi_Convert(t *testing.T) {
	Convey("Test: string is converted to number", t, func() {
		So(getConverted(""), ShouldEqual, 0)
		So(getConverted("1"), ShouldEqual, 1)
		So(getConverted("-1"), ShouldEqual, -1)
		So(getConverted("    1"), ShouldEqual, 1)
		So(getConverted("words    1"), ShouldEqual, 0)
		So(getConverted("-91283472332"), ShouldEqual, -2147483647)

	})
}

type (
	StringToAtoi interface {
		Convert(s string) int
	}
	stringToAtoi struct {
	}
)

func NewStringToAtoi() StringToAtoi {
	return stringToAtoi{}
}

func (stringToAtoi) Convert(s string) int {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return 0
	}
	isNeg := false
	if rune(s[0]) == '-' {
		isNeg = true
	}
	if !unicode.IsNumber(rune(s[0])) && rune(s[0]) != '-' {
		return 0
	}
	filtered := ""
	for _, val := range s {
		if unicode.IsNumber(val) {
			filtered += string(val)
		}
	}
	val, _ := strconv.ParseInt(filtered, 10, 32)
	result := int(val)

	if isNeg {
		result *= -1
	}
	if result < math.MinInt32 {
		return math.MinInt32
	}

	if result > math.MaxInt32 {
		return math.MaxInt32
	}
	return result
}
