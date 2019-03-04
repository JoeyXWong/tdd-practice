package strstr

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func findIndexOf(input string, target string) int {
	return NewStrStr().IndexOf(input, target)
}

func TestStrStr_IndexOf(t *testing.T) {
	Convey("Test: strstr should find the index of substring", t, func() {
		So(findIndexOf("", "1"), ShouldEqual, -1)
		So(findIndexOf("a", ""), ShouldEqual, 0)
		So(findIndexOf("a", "a"), ShouldEqual, 0)
		So(findIndexOf("hello", "l"), ShouldEqual, 2)
		So(findIndexOf("aaaaa", "ab"), ShouldEqual, -1)
	})
}
