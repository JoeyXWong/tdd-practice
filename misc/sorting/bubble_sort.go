package sorting

/**
Best case: O(n) //if already sorted
Worst case: O(n^2)
stable
in place
*/
type (
	bubbleSort struct {
	}

	BubbleSort interface {
		Sort(nums []int)
	}
)

func NewBubbleSort() BubbleSort {
	return bubbleSort{}

}

func (bubbleSort) Sort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		sorted := true
		j := len(nums) - 1
		for j > i {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				sorted = false
			}
			j--
		}
		if sorted {
			return
		}
	}
}
