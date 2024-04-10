package subarraysumequalsk

func subarraySum(nums []int, k int) int {
	m := make(map[int]int)
	pre := 0
	ans := 0
	m[0] = 1
	for _, num := range nums {
		pre += num
		ans += m[pre-k]
		m[pre] += 1
	}

	return ans
}

func subarraySumV1(nums []int, k int) int {
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}

	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if i == 0 && sums[j] == k {
				ans++
			}
			if i > 0 && sums[j]-sums[i-1] == k {
				ans++
			}
		}
	}

	return ans
}
