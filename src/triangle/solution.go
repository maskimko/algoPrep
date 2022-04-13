package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	row, col int
}

func dfs(triangle [][]int, memo map[coordinate]int, row, col int) int {
	if row >= len(triangle) {
		return 0 // This value will not change the total sum
	}
	if col >= len(triangle) {
		return 1<<31 - 1 // This value will not be in use anyway
	}
	if res, ok := memo[coordinate{row, col}]; ok {
		return res
	}
	min := 1<<31 - 1
	for i := 0; i < 2; i++ {
		currentMin := dfs(triangle, memo, row+1, col+i)
		if currentMin < min {
			min = currentMin
		}
	}
	memo[coordinate{row, col}] = min + triangle[row][col]
	return min + triangle[row][col]
}

func minimumTotal(triangle [][]int) int {
	memo := make(map[coordinate]int)
	return dfs(triangle, memo, 0, 0)
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
	triangleLength, _ := strconv.Atoi(scanner.Text())
	triangle := [][]int{}
	for i := 0; i < triangleLength; i++ {
		scanner.Scan()
		triangle = append(triangle, arrayAtoi(splitWords(scanner.Text())))
	}
	res := minimumTotal(triangle)
	fmt.Println(res)
}
