package reverse_integer

import "math"

type ReverseInteger interface {
	Reverse(x int) int
}

type reverseInteger struct {
}

func NewReverseInteger() ReverseInteger {
	return &reverseInteger{}
}

func (this *reverseInteger) Reverse(x int) int {
	result := 0
	original := x

	x = this.handleNegative(original, x)

	for x > 0 {
		result *= 10
		if this.isOverFlow(result, x) {
			return 0
		}
		result += x % 10
		x /= 10
	}

	return this.handleNegative(original, result)
}

func (this *reverseInteger) isOverFlow(result int, x int) bool {
	return result+(x%10) > math.MaxInt32
}

func (this *reverseInteger) handleNegative(preserved int, x int) int {
	if preserved < 0 {
		x *= -1
	}
	return x
}
