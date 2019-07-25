package sorting

import (
	"math/rand"
	"time"
)

type (
	QuickSortHoares interface {
		Sort(nums []int)
	}

	quickSortHoares struct {
	}
)

func NewQuickSortHoares() QuickSortHoares {
	return quickSortHoares{}
}

func (this quickSortHoares) Sort(nums []int) {
	if len(nums) < 2 {
		return
	}
	this.quickSortHelper(nums, 0, len(nums)-1)
}

func (this quickSortHoares) quickSortHelper(nums []int, left, right int) {
	if left >= right {
		return
	}

	rand.Seed(time.Now().UnixNano())
	pivotIndex := rand.Intn(right-left) + left
	this.swap(nums, left, pivotIndex)
	pn := nums[left]
	start := left + 1
	end := right
	for start <= right {
		for nums[start] < pn && start < end {
			start++
		}

		if start == end {
			break
		}

		for nums[end] > pn && end > start {
			end--
		}
		this.swap(nums, start, end)
	}
	this.swap(nums, left, end-1)

	quickSortHelper(nums, end, right)
	quickSortHelper(nums, left, end-2)
}

func (quickSortHoares) swap(nums []int, num1, num2 int) {
	nums[num1], nums[num2] = nums[num2], nums[num1]
}
