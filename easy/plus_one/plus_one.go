package plus_one

type (
	PlusOne interface {
		PlusOne(nums []int) []int
	}

	plusOne struct {
	}
)

func NewPlusOne() PlusOne {
	return plusOne{}
}

func (this plusOne) PlusOne(nums []int) []int {
	for i := len(nums) - 1; i >= 0; i-- {
		if !this.shouldCarryOver(nums[i]) {
			nums[i]++
			return nums
		}
		nums[i] = 0
	}
	return append([]int{1}, nums...)
}

func (this plusOne) shouldCarryOver(value int) bool {
	return value == 9
}
