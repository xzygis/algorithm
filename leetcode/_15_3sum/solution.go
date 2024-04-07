package sum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	var res [][]int
	for first := 0; first < len(nums)-2; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		for second := first + 1; second < len(nums)-1; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			third := len(nums) - 1
			for second < third && nums[first]+nums[second]+nums[third] > 0 {
				third--
			}

			if second == third {
				break
			}

			if nums[first]+nums[second]+nums[third] == 0 {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}

	return res
}
