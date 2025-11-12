package _347_top_k_frequent_elements

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	var heap []int
	for num := range freq {
		if len(heap) < k {
			heap = append(heap, num)
			siftUp(heap, freq, len(heap)-1)
		} else {
			if freq[num] > freq[heap[0]] {
				heap[0] = num
				siftDown(heap, freq, 0, len(heap))
			}
		}
	}

	return heap
}

func siftUp(nums []int, freq map[int]int, index int) {
	parent := (index - 1) / 2
	if parent >= 0 && freq[nums[index]] < freq[nums[parent]] {
		nums[parent], nums[index] = nums[index], nums[parent]
		siftUp(nums, freq, parent)
	}
}

func siftDown(nums []int, freq map[int]int, root int, heapSize int) {
	l, r, smallest := 2*root+1, 2*root+2, root
	if l < heapSize && freq[nums[l]] < freq[nums[smallest]] {
		smallest = l
	}

	if r < heapSize && freq[nums[r]] < freq[nums[smallest]] {
		smallest = r
	}

	if smallest != root {
		nums[smallest], nums[root] = nums[root], nums[smallest]
		siftDown(nums, freq, smallest, heapSize)
	}
}

func topKFrequentV1(nums []int, k int) []int {
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}

	var uniqNums []int
	for num := range countMap {
		uniqNums = append(uniqNums, num)
	}

	for i := len(uniqNums)/2 - 1; i >= 0; i-- {
		maxHeapify(uniqNums, countMap, i, len(uniqNums))
	}

	var ans []int
	for i := 0; i < k; i++ {
		ans = append(ans, uniqNums[0])
		uniqNums[0], uniqNums[len(uniqNums)-i-1] = uniqNums[len(uniqNums)-i-1], uniqNums[0]
		maxHeapify(uniqNums, countMap, 0, len(uniqNums)-i-1)
	}
	return ans
}

func maxHeapify(nums []int, countMap map[int]int, root int, heapSize int) {
	l, r, largest := 2*root+1, 2*root+2, root
	if l < heapSize && countMap[nums[l]] > countMap[nums[largest]] {
		largest = l
	}

	if r < heapSize && countMap[nums[r]] > countMap[nums[largest]] {
		largest = r
	}

	if largest != root {
		nums[largest], nums[root] = nums[root], nums[largest]
		maxHeapify(nums, countMap, largest, heapSize)
	}
}
