package plus_one

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func addOne(input []int) []int {
	return NewPlusOne().PlusOne(input)
}

func TestPlusOne_PlusOne(t *testing.T) {
	Convey("Test 0 value array is returned with array of 1", t, func() {
		result := addOne([]int{0})

		So(result, ShouldHaveLength, 1)
		So(result[0], ShouldEqual, 1)
	})

	Convey("Test value 9 should return two digit array", t, func() {
		result := addOne([]int{9})

		So(result, ShouldHaveLength, 2)
		So(result, ShouldResemble, []int{1, 0})
	})

	Convey("Test all 9's will return array length of one size larger", t, func() {
		result := addOne([]int{9, 9, 9, 9, 9})

		So(result, ShouldHaveLength, 6)
		So(result, ShouldResemble, []int{1, 0, 0, 0, 0, 0})
	})
}
