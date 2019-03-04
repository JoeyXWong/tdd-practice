package count_and_say

import (
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
)

func getCountAndSay(input int) string {
	return NewCountAndSay().Generate(input)
}

func TestCountAndSay_Generate(t *testing.T) {
	Convey("Test generate count and say", t, func() {
		So(getCountAndSay(1), ShouldEqual, "1")
		So(getCountAndSay(2), ShouldEqual, "11")
		So(getCountAndSay(3), ShouldEqual, "21")
	})
}

type countAndSay struct {
}

type CountAndSay interface {
	Generate(n int) string
}

func NewCountAndSay() CountAndSay {
	return countAndSay{}
}

func (countAndSay) Generate(n int) string {
	generated := "1"
	if n > 1 {
		temp := ""
		prev := generated[0]
		for i := range generated {
			count := 0
			val := generated[i]
			if prev == generated[i] {
				count++
				prev = val
			} else {
				output := strconv.Itoa(count) + string(prev)
				temp += output
				break
			}

		}
		generated = temp
	}
	return generated
}
