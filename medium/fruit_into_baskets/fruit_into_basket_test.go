package fruit_into_baskets

import (
	."github.com/smartystreets/goconvey/convey"
	"testing"
)

func harvestFruits(input []int) int {
	return NewFruitIntoBasket().GetFruits(input)
}

func TestFruitIntoBasket_GetFruits(t *testing.T){
	Convey("Test: 0-2 fruits should return 0-2 yields", t, func(){
		So(harvestFruits([]int{}), ShouldEqual, 0)
		So(harvestFruits([]int{1}), ShouldEqual, 1)
		So(harvestFruits([]int{1,2}), ShouldEqual, 2)
		So(harvestFruits([]int{1,2,1}), ShouldEqual, 3)
		So(harvestFruits([]int{1,2,3}), ShouldEqual, 2)
		So(harvestFruits([]int{0,1,2,2}), ShouldEqual, 3)
		//So(harvestFruits([]int{1,2,3,2,2}), ShouldEqual, 4)
		//So(harvestFruits([]int{3,3,3,1,2,1,1,2,3,3,4}), ShouldEqual, 5)
	})
}


