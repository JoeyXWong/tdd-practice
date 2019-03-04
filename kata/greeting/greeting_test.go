package greeting

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func greetName(input interface{}) string {
	return NewGreeting().Greet(input)
}

func TestGreeting_Greet(t *testing.T) {
	Convey("Test", t, func() {
		So(greetName("Bob"), ShouldEqual, "Hello, Bob.")
		So(greetName(""), ShouldEqual, "Hello, my friend.")
		So(greetName("BOB"), ShouldEqual, "HELLO, BOB.")
		So(greetName([]string{"Jack", "Jill"}), ShouldEqual, "Hello, Jack and Jill.")
	})
}
