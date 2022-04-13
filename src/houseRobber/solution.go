package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func rob(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i < 2 {
			dp[i] = nums[i]
			continue
		}
		if dp[i-1] > nums[i]+dp[i-2] {
			dp[i] = dp[i-1]
		} else {
			dp[i] = nums[i] + dp[i-2]
		}
	}
	return dp[len(dp)-1]
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
	res := rob(nums)
	fmt.Println(res)
}
