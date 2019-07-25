package sorting

/**
Best case: O(n^2)
Worst case: O(n^2)
Not stable
in place
*/
type (
	selectionSort struct {
	}

	SelectionSort interface {
		Sort(nums []int)
	}
)

func NewSelectionSort() SelectionSort {
	return selectionSort{}
}

func (this selectionSort) Sort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[minIndex] > nums[j] {
				minIndex = j
			}
		}
		if minIndex != i {
			this.swap(nums, minIndex, i)
		}
	}
}

func (selectionSort) swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
