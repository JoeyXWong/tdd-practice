package rotate_image

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func createMatrix(input [][]int) [][]int {
	return NewRotateImage().Rotate(input)

}

func TestRotateImage_Rotate(t *testing.T) {

	Convey("Test 1 by 1 matrix is rotated", t, func() {
		So(createMatrix([][]int{{1}}), ShouldResemble, [][]int{{1}})
	})

	Convey("Test 2 by 2 matrix is rotated", t, func() {
		result := createMatrix([][]int{
			{1, 2},
			{3, 4},
		})
		expected := [][]int{
			{3, 1},
			{4, 2},
		}

		So(result, ShouldResemble, expected)
	})

	Convey("Test 3 by 3 matrix is rotated", t, func() {
		result := createMatrix([][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		})
		expected := [][]int{
			{7, 4, 1},
			{8, 5, 2},
			{9, 6, 3},
		}

		So(result, ShouldResemble, expected)
	})

	Convey("Test 4 by 4 matrix is rotated", t, func() {
		result := createMatrix([][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		})
		expected := [][]int{
			{13, 9, 5, 1},
			{14, 10, 6, 2},
			{15, 11, 7, 3},
			{16, 12, 8, 4},
		}

		So(result, ShouldResemble, expected)
	})
}
