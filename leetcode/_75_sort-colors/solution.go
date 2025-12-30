package _75_sort_colors

func sortColors(nums []int) {
	left, mid, right := 0, 0, len(nums)-1
	for mid <= right {
		if nums[mid] == 0 {
			nums[left], nums[mid] = nums[mid], nums[left]
			left++
			mid++
		} else if nums[mid] == 2 {
			nums[mid], nums[right] = nums[right], nums[mid]
			right--
			//该分支不能确定交换过来的right位置的值是什么，所以需要下一轮继续检查
		} else {
			mid++
		}
	}
}

func sortColorsV1(nums []int) {
	pos := 0
	for i := pos; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[pos], nums[i] = nums[i], nums[pos]
			pos++
		}
	}

	for i := pos; i < len(nums); i++ {
		if nums[i] == 1 {
			nums[pos], nums[i] = nums[i], nums[pos]
			pos++
		}
	}
}
