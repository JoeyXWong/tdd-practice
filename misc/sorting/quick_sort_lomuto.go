package sorting

import (
	"math/rand"
	"time"
)

type (
	QuickSortLomuto interface {
		Sort(nums []int)
	}

	quickSortLomuto struct {
	}
)

func NewQuickSortLomuto() QuickSortLomuto {
	return quickSortLomuto{}
}

func (quickSortLomuto) Sort(nums []int) {
	quickSortHelper(nums, 0, len(nums)-1)
}

func quickSortHelper(nums []int, left int, right int) {
	if left >= right {
		return
	}

	rand.Seed(time.Now().UnixNano())
	pivotIndex := rand.Intn(right-left) + left
	pivotNumber := nums[pivotIndex]
	nums[left], nums[pivotIndex] = nums[pivotIndex], nums[left]

	lowerPtr := left + 1
	higherPtr := left + 1

	for lowerPtr <= right && nums[lowerPtr] < pivotNumber {
		lowerPtr++
	}

	higherPtr = lowerPtr
	for higherPtr <= right {
		if nums[higherPtr] < pivotNumber {
			nums[lowerPtr], nums[higherPtr] = nums[higherPtr], nums[lowerPtr]
			lowerPtr++
		}
		higherPtr++
	}

	nums[left], nums[lowerPtr-1] = nums[lowerPtr-1], nums[left]

	quickSortHelper(nums, left, lowerPtr-1)
	quickSortHelper(nums, lowerPtr, right)
}
