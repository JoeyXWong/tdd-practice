package single_number

import (
	."github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSingleNumber_GetSingleNumber(t *testing.T){
	Convey("Test: Duplicate number is returned from the list", t, func(){
		input := []int{1,2,1,4,2, 4, 3}
		duplicate := NewSingleNumber()
		result := duplicate.GetSingleNumber(input)

		So(result, ShouldEqual, 3)
	})
}


