package _189_rotate_array

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, left int, right int) {
	for left < right {
		temp := nums[left]
		nums[left] = nums[right]
		nums[right] = temp
		left++
		right--
	}
}
