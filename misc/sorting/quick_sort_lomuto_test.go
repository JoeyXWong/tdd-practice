package sorting

import (
	. "github.com/smartystreets/goconvey/convey"

	"testing"
)

func TestQuickSortLomuto_Sort(t *testing.T) {
	Convey("Test quick Sort, no duplicates", t, func() {
		input := []int{3, 4, 1, 32, 11, 100, 6, 2, 8, 10, 22}
		expectedResult := []int{1, 2, 3, 4, 6, 8, 10, 11, 22, 32, 100}
		sorter := NewQuickSortLomuto()
		sorter.Sort(input)

		So(input, ShouldResemble, expectedResult)
	})

	Convey("Test quick Sort, no duplicates", t, func() {
		input := []int{3, 4, 1, 2, 2, 3, 10}
		expectedResult := []int{1, 2, 2, 3, 3, 4, 10}
		sorter := NewQuickSortLomuto()
		sorter.Sort(input)

		So(input, ShouldResemble, expectedResult)
	})
}
