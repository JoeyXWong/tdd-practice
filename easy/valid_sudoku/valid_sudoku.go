package valid_sudoku

type (
	ValidSudoku interface {
		IsValid(sudoku [][]string) bool
	}

	validSudoku struct {
	}

	point struct {
		x int
		y int
	}
)

func NewValidSudoku() ValidSudoku {
	return validSudoku{}
}

func (validSudoku) IsValid(sudoku [][]string) bool {
	populateMap := make(map[point]map[string]bool)

	return true
}
