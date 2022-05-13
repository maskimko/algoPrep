package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
NOT SOLVED

Check https://www.geeksforgeeks.org/palindrome-partitioning-dp-17/
*/

func partition2(s string) int {
	dp := make([][]bool, len(s), len(s))
	partitionCount := make([]int, len(s), len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s), len(s))
		for j := 0; j <= i; j++ {
			dp[i][j] = true
		}
	}
	chars := []rune(s)
	for l := 1; l < len(s); l++ {
		for i := 0; i < len(s)-l; i++ {
			j := i + l
			one := chars[i]
			other := chars[j]
			dp[i][j] = dp[i+1][j-1] && one == other
		}
	}
	partitionCount[0] = 1
	for i := 1; i < len(s); i++ {
		for j := len(s) - 1; j >= len(s)-i-1; j-- {
			if dp[i][j] {
				partitionCount[i]++
			}
		}
	}
	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	res := partition2(s)
	fmt.Println(res)
}
