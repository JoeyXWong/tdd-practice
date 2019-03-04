package single_number

type (
	SingleNumber interface {
		GetSingleNumber(nums []int) int
	}

	singleNumber struct {
	}
)

func NewSingleNumber() SingleNumber {
	return singleNumber{}
}

func (singleNumber) GetSingleNumber(nums []int) int {
	duplicate := 0
	for _, value := range nums{
		duplicate ^= value
	}
	return duplicate
}
