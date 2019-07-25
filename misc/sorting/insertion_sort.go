package sorting

/**
Best case: O(n)
Worst case: O(n^2)
stable
in place
*/
type (
	InsertionSort interface {
		Sort(nums []int)
	}

	insertionSort struct {
	}
)

func NewInsertionSort() InsertionSort {
	return insertionSort{}
}

func (insertionSort) Sort(nums []int) {
	for i := 1; i < len(nums); i++ {
		j := i - 1
		valueToSort := nums[i]

		for j >= 0 && nums[j] > valueToSort {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = valueToSort
	}
}
