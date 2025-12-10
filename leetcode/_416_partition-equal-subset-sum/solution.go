package _416_partition_equal_subset_sum

func canPartition(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}

	return dp[target]
}
