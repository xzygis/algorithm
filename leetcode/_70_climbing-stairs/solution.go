package _70_climbing_stairs

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	n1 := 1
	n2 := 2
	for i := 3; i <= n; i++ {
		n3 := n1 + n2
		n1 = n2
		n2 = n3
	}

	return n2
}

func climbStairsV1(n int) int {
	if n <= 2 {
		return n
	}

	return climbStairs(n-1) + climbStairs(n-2)
}
