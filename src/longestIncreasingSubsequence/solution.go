package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func longestSubLen(nums []int) int {
	dp := make([]int, len(nums), len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	best := 0
	for i := 1; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
					if dp[i] > best {
						best = dp[i]
					}
				}
			}
		}
	}
	return best
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
	nums := arrayAtoi(splitWords(scanner.Text()))
	res := longestSubLen(nums)
	fmt.Println(res)
}
