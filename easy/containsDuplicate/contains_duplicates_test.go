package containsDuplicate

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)


func assertDuplicate(input []int) bool {
	checker := NewChecker()
	return checker.ContainsDuplicate(input)
}

func TestContainsDuplicate_ContainsDuplicate(t *testing.T){
	Convey("Test should return false there are duplicates in array", t, func(){
		So(assertDuplicate([]int{1, 2, 3}), ShouldBeFalse)
		So(assertDuplicate([]int{}), ShouldBeFalse)
	})

	Convey("Test should return true if one number occurs twice", t, func() {
		So(assertDuplicate([]int{1,1}), ShouldBeTrue)
	})
}


