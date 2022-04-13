package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid), len(grid))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(grid[i]), len(grid[i]))
	}
	for i := 0; i < len(dp); i++ {
		for h := 0; h < len(dp[i]); h++ {
			if i == 0 && h == 0 {
				dp[i][h] = grid[i][h]
				continue
			}
			sumTop := 1<<31 - 1
			sumLeft := 1<<31 - 1
			if i-1 >= 0 {
				sumTop = dp[i-1][h] + grid[i][h]
			}
			if h-1 >= 0 {
				sumLeft = dp[i][h-1] + grid[i][h]
			}
			if sumTop < sumLeft {
				dp[i][h] = sumTop
			} else {
				dp[i][h] = sumLeft
			}
		}
	}
	return dp[len(dp)-1][len(dp[len(dp)-1])-1]
}

func arrayAtoi(arr []string) []int {
	res := []int{}
	for _, x := range arr {
		v, _ := strconv.Atoi(x)
		res = append(res, v)
	}
	return res
}

func splitWords(s string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, " ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	gridLength, _ := strconv.Atoi(scanner.Text())
	grid := [][]int{}
	for i := 0; i < gridLength; i++ {
		scanner.Scan()
		grid = append(grid, arrayAtoi(splitWords(scanner.Text())))
	}
	res := minPathSum(grid)
	fmt.Println(res)
}
