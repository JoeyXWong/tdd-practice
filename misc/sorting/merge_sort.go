package sorting

type (
	mergeSort struct {
	}

	MergeSort interface {
		Sort(nums []int)
	}
)

func NewMergeSort() MergeSort {
	return mergeSort{}

}

func (mergeSort) Sort(nums []int) {

	mid := len(nums) / 2

	l := mergeSortHelper(nums[:mid])
	r := mergeSortHelper(nums[mid:])

	aux := merge(l, r)
	copy(nums, aux)
}

func mergeSortHelper(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	mid := len(nums) / 2

	l := mergeSortHelper(nums[:mid])
	r := mergeSortHelper(nums[mid:])

	return merge(l, r)
}

func merge(left []int, right []int) []int {

	aux := make([]int, len(right)+len(left))
	li := 0
	ri := 0
	mi := 0

	for li < len(left) && ri < len(right) {
		if left[li] <= right[ri] {
			aux[mi] = left[li]
			li++
		} else {
			aux[mi] = right[ri]
			ri++
		}
		mi++
	}

	for li < len(left) {
		aux[mi] = left[li]
		li++
		mi++
	}

	for ri < len(right) {
		aux[mi] = right[ri]
		ri++
		mi++
	}

	return aux
}
