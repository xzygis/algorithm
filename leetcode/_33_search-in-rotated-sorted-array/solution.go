package _33_search_in_rotated_sorted_array

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] { //左侧有序。这里必须加等号，对于下表left=mid的场景，应该归结为左侧有序。
			if nums[left] <= target && target < nums[mid] { //target在[left, mid)之间
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { //右侧有序
			if nums[mid] < target && target <= nums[right] { //target在(mid, right]之间
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
