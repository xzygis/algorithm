package _238_product_of_array_except_self

func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	left[0] = nums[0]
	right := make([]int, len(nums))
	right[len(nums)-1] = nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i]
	}

	ans := make([]int, len(nums))
	ans[0] = right[1]
	ans[len(nums)-1] = left[len(nums)-2]
	for i := 1; i < len(nums)-1; i++ {
		ans[i] = left[i-1] * right[i+1]
	}

	return ans
}
