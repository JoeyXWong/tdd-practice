package valid_palindrome

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
	"unicode"
)

func isPalindrome(input string) bool {
	return NewValidPalindrome().IsValid(input)
}

func TestValidPalindrome_IsValid(t *testing.T) {
	Convey("Test: Palindrome is valid for empty case", t, func() {
		So(isPalindrome(""), ShouldEqual, true)
		So(isPalindrome("a"), ShouldEqual, true)
		So(isPalindrome("ab"), ShouldEqual, false)
		So(isPalindrome("aba"), ShouldEqual, true)
		So(isPalindrome("abba"), ShouldEqual, true)
		So(isPalindrome("Ababa"), ShouldEqual, true)
		So(isPalindrome("adaba"), ShouldEqual, false)
		So(isPalindrome("a baba"), ShouldEqual, true)
		So(isPalindrome("a :baba"), ShouldEqual, true)
		So(isPalindrome("A man, a plan, a canal: Panama"), ShouldEqual, true)
		So(isPalindrome(",."), ShouldEqual, true)
		So(isPalindrome("0P"), ShouldEqual, false)
	})
}

type (
	ValidPalindrome interface {
		IsValid(s string) bool
	}

	validPalindrome struct {
	}
)

func NewValidPalindrome() ValidPalindrome {
	return &validPalindrome{}
}

func (this *validPalindrome) IsValid(s string) bool {
	if len(s) >= 2 {
		s = strings.ToLower(s)
		i := 0
		for !IsAlphaNumeric(s[i]) && len(s)-1 > i {
			i++
		}
		j := len(s) - 1
		for !IsAlphaNumeric(s[j]) && j > 0 {
			j--
		}
		if i < j {
			return s[i] != s[j] && this.IsValid(s[i+1:j])
		}
	}
	return true
}

func IsAlphaNumeric(c uint8) bool {
	return (c >= 48 && c <= 57) || unicode.IsLetter(rune(c))
}
