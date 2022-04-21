package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
NOT SOLVED
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
	for l := 1; l < len(s); l++ {
		for i := 0; i < len(s)-l; i++ {
			j := i + l
			dp[i][j] = dp[i+1][j-1] && []rune(s)[i] == []rune(s)[j]
		}
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
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
