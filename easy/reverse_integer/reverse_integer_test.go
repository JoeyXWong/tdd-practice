package reverse_integer

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func getReverse(input int) int {
	return NewReverseInteger().Reverse(input)
}

func TestReverseInteger_Reverse(t *testing.T) {
	Convey("Test single digit digit is reversed", t, func() {
		So(getReverse(1), ShouldEqual, 1)
	})

	Convey("Test same numbers should return the same", t, func() {
		So(getReverse(22), ShouldEqual, 22)
	})

	Convey("Test random numbers should return reversed", t, func() {
		So(getReverse(23513), ShouldEqual, 31532)
	})

	Convey("Test overflow should return 0", t, func() {
		So(getReverse(214748364214), ShouldEqual, 0)
	})

	Convey("Test should reverse negative integers", t, func() {
		So(getReverse(-23513), ShouldEqual, -31532)
	})
}
