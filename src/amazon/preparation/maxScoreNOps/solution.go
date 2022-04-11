package maxScoreNOps

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func getScore(i, x, y int, nums []int) int {
	return i * gcd(nums[x], nums[y])
}

func dfs(mask int, nums []int, memo map[int]int, numOfOp int) int {

	if mask == 0 {
		return 0
	}
	if score, ok := memo[mask]; ok {
		return score
	}
	max := 0
	for i := 0; i < len(nums)-1; i++ {
		if mask&(1<<i) == 0 {
			continue
		}
		for k := i + 1; k < len(nums); k++ {
			if mask&(1<<k) == 0 {
				continue
			}
			score := getScore(numOfOp, i, k, nums)
			m := mask & ^(1 << i) & ^(1 << k)
			rest := dfs(m, nums, memo, numOfOp+1)
			if max < score+rest {
				max = score + rest
			}
			if _, ok := memo[mask]; !ok {
				memo[mask] = max
			}
		}
	}
	return max
}

func maxScore(nums []int) int {
	memo := make(map[int]int)
	return dfs(1<<len(nums)-1, nums, memo, 1)
}
