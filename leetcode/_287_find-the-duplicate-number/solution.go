package _287_find_the_duplicate_number

func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]
	//找到相遇点
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	//找到环的入口
	fast = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}

func findDuplicateV1(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		count := 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}

		if count > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
