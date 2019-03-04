package intersection_two_array_2

import (
	."github.com/smartystreets/goconvey/convey"
	"testing"
)

func getIntersection(inputOne []int, inputTwo []int) []int {
	return NewIntersectionOfTwoArray().GetIntersect(inputOne, inputTwo)
}

func TestIntersectionOfTwoArray_GetIntersect(t *testing.T){
	Convey("Test common elements from the two arrays will be returned", t, func(){
		result := getIntersection( []int{1,2,3,2}, []int{2,3,4,5,2})

		So(result, ShouldHaveLength, 3)
		So(result, ShouldResemble, []int{2,3,2})
	})

	Convey("Test non match should return an empty array", t, func(){
		result := getIntersection([]int{}, []int{1})

		So(result, ShouldHaveLength, 0)
	})
}

