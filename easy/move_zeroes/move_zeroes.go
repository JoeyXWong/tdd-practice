package move_zeroes

type (
	MoveZeroes interface {
		MoveZeroes(nums []int) []int
	}

	moveZeroes struct {
	}
)

func NewMoveZeroes() MoveZeroes {
	return moveZeroes{}
}

func (moveZeroes) MoveZeroes(nums []int) []int {
	populatedIdx := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[populatedIdx], nums[i] = nums[i], nums[populatedIdx]
			populatedIdx++
		}
	}
	return nums
}
