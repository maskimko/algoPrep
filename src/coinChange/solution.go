package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func minCoin(coins []int, dp []int, c int) (int, int) {
	inf := 1<<31 - 1
	minimum := inf
	coinIdx := -1
	prev := -1
	for i := 0; i < len(coins); i++ {
		val := c - coins[i]
		if val < 0 {
			continue
		}
		if dp[val] < minimum {
			coinIdx = i
			prev = val
			minimum = dp[val]
		}
	}
	return coinIdx, prev
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		if i == 0 {
			dp[0] = 0
			continue
		}
		cIdx, prev := minCoin(coins, dp, i)
		if cIdx >= 0 && dp[prev] >= 0 {
			dp[i] = 1 + dp[prev]
		} else {
			dp[i] = -1
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
	coins := arrayAtoi(splitWords(scanner.Text()))
	scanner.Scan()
	amount, _ := strconv.Atoi(scanner.Text())
	res := coinChange(coins, amount)
	fmt.Println(res)
}
