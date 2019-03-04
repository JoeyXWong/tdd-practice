package reverse_string

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func getReversed(input []byte) []byte {
	return NewReverseString().Reverse(input)
}

func TestReverseString_Reverse(t *testing.T) {
	Convey("Test empty array is returned as empty", t, func() {
		So(getReversed([]byte{}), ShouldHaveLength, 0)
	})

	Convey("Test single letter array returns the same array", t, func() {
		So(getReversed([]byte{'a'}), ShouldResemble, []byte{'a'})
	})

	Convey("Test array is returned as reversed", t, func() {
		So(getReversed([]byte{'a', 'b', 'b', 'c'}), ShouldResemble, []byte{'c', 'b', 'b', 'a'})
	})
}
