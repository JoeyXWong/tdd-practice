package sorting

/**
0 index parent (i-1)/2
1 index parent i/2
0 index left child (i * 2) + 1
0 index right child (i*2) + 2
1 index left child i*2
1 index right child (i*2) + 1
*/
type heapSort struct{}

type HeapSort interface {
	Sort(nums []int)
}

func NewHeapSort() HeapSort {

	return &heapSort{}
}

func (this heapSort) Sort(nums []int) {
	this.buildHeap(nums)
	this.heapSort(nums)
}

func (this heapSort) buildHeap(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		curr := i
		parent := (curr - 1) / 2
		for parent >= 0 && nums[parent] < nums[curr] {
			this.swap(nums, parent, curr)
			curr = parent
			parent = (curr - 1) / 2
		}
	}
}

func (this heapSort) swap(nums []int, num1 int, num2 int) {
	nums[num1], nums[num2] = nums[num2], nums[num1]
}

func (this heapSort) heapSort(nums []int) {
	sortIdx := len(nums) - 1

	for sortIdx > 0 {
		this.swap(nums, sortIdx, 0)
		this.bubbleDown(nums, sortIdx)
		sortIdx--
	}
}

func (this heapSort) bubbleDown(nums []int, sortedIdx int) {

	parent := 0
	child1 := (parent * 2) + 1
	child2 := child1 + 1

	for child1 < sortedIdx && child2 < sortedIdx {
		if nums[parent] > nums[child1] && nums[parent] > nums[child2] {
			return
		}

		if nums[child1] > nums[child2] {
			this.swap(nums, child1, parent)
			parent = child1
		} else {
			this.swap(nums, child2, parent)
			parent = child2
		}

		child1 = (parent * 2) + 1
		child2 = child1 + 1
	}

	if child1 < sortedIdx && nums[child1] > nums[parent] {
		this.swap(nums, child1, parent)
	}
}
