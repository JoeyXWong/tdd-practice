package sorting

import (
	. "github.com/smartystreets/goconvey/convey"

	"testing"
)

func TestHeapSort_Sort(t *testing.T) {
	Convey("Test Heap Sort, no duplicates", t, func() {
		input := []int{3, 4, 1, 6, 2, 8, 10}
		expectedResult := []int{1, 2, 3, 4, 6, 8, 10}
		sorter := NewHeapSort()
		sorter.Sort(input)

		So(input, ShouldResemble, expectedResult)
	})

	Convey("Test Heap Sort, no duplicates", t, func() {
		input := []int{3, 4, 1, 2, 2, 3, 10}
		expectedResult := []int{1, 2, 2, 3, 3, 4, 10}
		sorter := NewHeapSort()
		sorter.Sort(input)

		So(input, ShouldResemble, expectedResult)
	})
}
