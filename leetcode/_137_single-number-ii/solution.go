package _137_single_number_ii

func singleNumber(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		sum := 0
		for _, num := range nums {
			sum += (num >> i) & 1
		}

		// 如果sum为3的倍数为1，说明3个相同的位为0或者1，而且结果的当前位为1
		if sum%3 == 1 {
			ans |= 1 << i
		}
	}

	return int(ans)
}

func singleNumberV1(nums []int) int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	for k, v := range freq {
		if v == 1 {
			return k
		}
	}

	return -1
}
