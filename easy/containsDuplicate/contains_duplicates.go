package containsDuplicate

type (
	Checker interface {
		ContainsDuplicate(ints []int) bool
	}

	checker struct {
	}
)

func NewChecker() Checker {
	return checker{}
}

// O(n) space and O(n) time
func (checker) ContainsDuplicate(nums []int) bool {
	duplicates := make(map[int]bool)
	for _, value := range nums {
		if duplicates[value] {
			return true
		}
		duplicates[value] = true
	}
	return false
}
