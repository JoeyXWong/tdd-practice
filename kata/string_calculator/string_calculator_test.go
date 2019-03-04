package string_calculator

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func calculateAdd(inputOne string, inputTwo string) int {
	return NewStringCalculator().Add(inputOne, inputTwo)
}

func TestStringCalculator_Add(t *testing.T) {
	Convey("Test: Can take 0 input and return 0", t, func() {
		So(calculateAdd("", ""), ShouldEqual, 0)
		So(calculateAdd("", "1"), ShouldEqual, 1)
		So(calculateAdd("1", ""), ShouldEqual, 1)
	})
}
