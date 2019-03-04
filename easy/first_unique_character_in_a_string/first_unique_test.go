package first_unique_character_in_a_string

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func findFirstUnique(input string) int {
	return NewFirstUniqueCharacter().FirstUnique(input)
}

func TestFirstUniqueCharacter_FirstUnique(t *testing.T) {
	Convey("Test returns -1 for empty string", t, func() {
		So(findFirstUnique(""), ShouldEqual, -1)
	})

	Convey("Test single letter should return 0", t, func() {
		So(findFirstUnique("a"), ShouldEqual, 0)
	})

	Convey("Test multiple uniques should return the first unique", t, func() {
		So(findFirstUnique("aadgfgs"), ShouldEqual, 2)
	})

	Convey("Test no uniques should return -1", t, func() {
		So(findFirstUnique("saadgdfgfs"), ShouldEqual, -1)
	})
}
