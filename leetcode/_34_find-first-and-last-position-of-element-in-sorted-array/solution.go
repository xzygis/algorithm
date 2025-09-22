package _34_find_first_and_last_position_of_element_in_sorted_array

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	leftBound := findFirst(nums, target)
	if leftBound == len(nums) || nums[leftBound] != target {
		//左边界不是target（没找到target）
		return []int{-1, -1}
	}

	rightBound := findLast(nums, target) - 1
	return []int{leftBound, rightBound}
}

// 第一个大于等于target的元素索引
func findFirst(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

// 第一个大于target的元素索引
func findLast(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
