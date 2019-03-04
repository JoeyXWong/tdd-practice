package valid_sudoku

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestValidSudoku_IsValid(t *testing.T) {
	Convey("Test valid sudoku should return true", t, func() {
		input := [][]string{
			{"5", "3", ".", ".", "7", ".", ".", ".", "."},
			{"6", ".", ".", "1", "9", "5", ".", ".", "."},
			{".", "9", "8", ".", ".", ".", ".", "6", "."},
			{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
			{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
			{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
			{".", "6", ".", ".", ".", ".", "2", "8", "."},
			{".", ".", ".", "4", "1", "9", ".", ".", "5"},
			{".", ".", ".", ".", "8", ".", ".", "7", "9"},
		}
		validator := NewValidSudoku()
		result := validator.IsValid(input)

		So(result, ShouldBeTrue)
	})

	Convey("Test invalid sudoku from rows", t, func() {
		input := [][]string{
			{"5", "3", ".", ".", "7", ".", ".", "3", "."},
			{"6", ".", ".", "1", "9", "5", ".", ".", "."},
			{".", "9", "8", ".", ".", ".", ".", "6", "."},
			{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
			{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
			{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
			{".", "6", ".", ".", ".", ".", "2", "8", "."},
			{".", ".", ".", "4", "1", "9", ".", ".", "5"},
			{".", ".", ".", ".", "8", ".", ".", "7", "9"},
		}

		validator := NewValidSudoku()
		result := validator.IsValid(input)

		So(result, ShouldBeFalse)
	})
}
