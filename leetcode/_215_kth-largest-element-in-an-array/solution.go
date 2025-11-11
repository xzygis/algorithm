package _215_kth_largest_element_in_an_array

func findKthLargest(nums []int, k int) int {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		maxHeapify(nums, i, len(nums))
	}

	heapSize := len(nums)
	// 第k-1大的元素应该是交换至位置len(nums)-(k-1)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}

	return nums[0]
}

func maxHeapify(nums []int, root int, heapSize int) {
	l, r, largest := 2*root+1, 2*root+2, root
	if l < heapSize && nums[l] > nums[largest] {
		largest = l
	}

	if r < heapSize && nums[r] > nums[largest] {
		largest = r
	}

	if largest != root {
		nums[root], nums[largest] = nums[largest], nums[root]
		maxHeapify(nums, largest, heapSize)
	}
}
