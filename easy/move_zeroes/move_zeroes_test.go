package move_zeroes

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMoveZeroes_MoveZeroes(t *testing.T) {
	Convey("Test empty array should return empty", t, func() {
		result := NewMoveZeroes().MoveZeroes([]int{})

		So(result, ShouldHaveLength, 0)
	})

	Convey("Test Zero at the end should return the same array", t, func() {
		result := NewMoveZeroes().MoveZeroes([]int{1, 0})

		So(result, ShouldHaveLength, 2)
		So(result, ShouldResemble, []int{1, 0})
	})

	Convey("Test One Zero in the beginning should be moved to the back of the array", t, func() {
		result := NewMoveZeroes().MoveZeroes([]int{0, 1})

		So(result, ShouldHaveLength, 2)
		So(result, ShouldResemble, []int{1, 0})
	})

	Convey("Test multiple zeroes should all be returned to the end with the preserved ordering of non zero elements", t, func() {
		result := NewMoveZeroes().MoveZeroes([]int{0, 1, 2, 0, 11, 0, 0, 0, 1, 2, 3, 6, 0})

		So(result, ShouldHaveLength, 13)
		So(result, ShouldResemble, []int{1, 2, 11, 1, 2, 3, 6, 0, 0, 0, 0, 0, 0})
	})
}
