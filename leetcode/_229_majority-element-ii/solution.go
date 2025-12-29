package _229_majority_element_ii

func majorityElement(nums []int) []int {
	element1 := 0
	vote1 := 0
	element2 := 0
	vote2 := 0
	for _, num := range nums {
		if vote1 > 0 && num == element1 {
			vote1++
		} else if vote2 > 0 && num == element2 {
			vote2++
		} else if vote1 == 0 {
			element1 = num
			vote1 = 1
		} else if vote2 == 0 {
			element2 = num
			vote2 = 1
		} else {
			vote1--
			vote2--
		}
	}

	count1, count2 := 0, 0
	for _, num := range nums {
		if num == element1 {
			count1++
		} else if num == element2 {
			count2++
		}
	}

	var ans []int
	if vote1 > 0 && count1 > len(nums)/3 {
		ans = append(ans, element1)
	}
	if vote2 > 0 && count2 > len(nums)/3 {
		ans = append(ans, element2)
	}

	return ans
}
