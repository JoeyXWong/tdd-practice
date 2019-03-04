package is_anagram

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func checkAnagram(input string, target string) bool {
	return NewIsAnagram().IsAnagram(input, target)
}

func TestIsAnagram_IsAnagram(t *testing.T) {

	Convey("Test empty string should return true", t, func() {
		So(checkAnagram("", ""), ShouldBeTrue)
	})

	Convey("Test one letter should return true", t, func() {
		So(checkAnagram("a", "a"), ShouldBeTrue)
	})

	Convey("Test valid anagram should return true", t, func() {
		So(checkAnagram("adeagd", "aaddeg"), ShouldBeTrue)
	})

	Convey("Test mismatch letters should return false", t, func() {
		So(checkAnagram("ase", "abf"), ShouldBeFalse)
	})

	Convey("Test is mismatch from target to input, return false", t, func() {
		So(checkAnagram("abgee", "abged"), ShouldBeFalse)
	})
}
