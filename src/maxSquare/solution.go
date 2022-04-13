package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func min(x, y, z int) int {
	r := x
	if y < r {
		r = y
	}
	if z < r {
		r = z
	}
	return r
}

func maximalSquare(matrix [][]int) int {
	dp := make([][]int, len(matrix), len(matrix))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[i]), len(matrix[i]))
	}
	maxSq := 0
	for i := 0; i < len(dp[0]); i++ {
		if matrix[0][i] == 1 {
			if maxSq < 1 {
				maxSq = 1
			}
			dp[0][i] = 1
		}
	}
	for i := 0; i < len(dp); i++ {
		if matrix[i][0] == 1 {
			if maxSq < 1 {
				maxSq = 1
			}
			dp[i][0] = 1
		}
	}
	for i := 1; i < len(dp); i++ {
		for k := 1; k < len(dp[i]); k++ {
			if matrix[i][k] == 1 {
				dp[i][k] = 1 + min(dp[i-1][k], dp[i][k-1], dp[i-1][k-1])
				if dp[i][k] > maxSq {
					maxSq = dp[i][k]
				}
			}
		}
	}
	return maxSq * maxSq
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

func readAndRun(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	matrixLength, _ := strconv.Atoi(scanner.Text())
	matrix := [][]int{}
	for i := 0; i < matrixLength; i++ {
		scanner.Scan()
		matrix = append(matrix, arrayAtoi(splitWords(scanner.Text())))
	}
	return maximalSquare(matrix)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	matrixLength, _ := strconv.Atoi(scanner.Text())
	matrix := [][]int{}
	for i := 0; i < matrixLength; i++ {
		scanner.Scan()
		matrix = append(matrix, arrayAtoi(splitWords(scanner.Text())))
	}
	res := maximalSquare(matrix)
	fmt.Println(res)
}
