package sorting

import (
	. "github.com/smartystreets/goconvey/convey"

	"testing"
)

func TestMergeSort_Sort(t *testing.T) {
	Convey("Test Merge Sort, no duplicates", t, func() {
		input := []int{3, 4, 1, 6, 2, 8, 10}
		expectedResult := []int{1, 2, 3, 4, 6, 8, 10}
		sorter := NewMergeSort()
		sorter.Sort(input)

		So(input, ShouldResemble, expectedResult)
	})

	Convey("Test Merge Sort, no duplicates", t, func() {
		input := []int{3, 4, 1, 2, 2, 3, 10}
		expectedResult := []int{1, 2, 2, 3, 3, 4, 10}
		sorter := NewMergeSort()
		sorter.Sort(input)

		So(input, ShouldResemble, expectedResult)
	})
}
