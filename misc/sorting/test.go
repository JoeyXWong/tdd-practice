package sorting

func Sort(nums []int) {
	start := 0

	for start < len(nums) {
		min := nums[start]
		idx := start
		for i := start; i < len(nums); i++ {
			if min > nums[i] {
				min = nums[i]
				idx = i
			}
		}
		nums[start], nums[idx] = nums[idx], nums[start]
		start++
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
