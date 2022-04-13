package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func uniquePaths(m int, n int) int {
	dp := make([][]int, n, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m, m)
	}
	dp[0][0] = 1
	for v := 0; v < len(dp); v++ {
		for h := 0; h < len(dp[v]); h++ {
			// top
			top := 0
			if v-1 >= 0 {
				top = dp[v-1][h]
			}
			// left
			left := 0
			if h-1 >= 0 {
				left = dp[v][h-1]
			}
			if dp[v][h] < top+left {
				dp[v][h] = top + left
			}
		}
	}
	return dp[len(dp)-1][len(dp[len(dp)-1])-1]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	res := uniquePaths(m, n)
	fmt.Println(res)
}
