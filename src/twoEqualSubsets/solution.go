package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
top to bottom DFS solution
*/

func dfs(nums []int, target, mask int, memo map[int]int) bool {
	if len(nums) < 2 {
		return false
	}
	if res, ok := memo[mask]; ok {
		return res == target
	}
	if target == 0 {
		return true
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		if mask&(1<<i) == 0 {
			continue
		}
		sum += nums[i]
	}
	memo[mask] = sum
	for i := 0; i < len(nums); i++ {
		if mask&(1<<i) == 0 {
			continue
		}
		m := mask & ^(1 << i)
		t := target - nums[i]
		ok := dfs(nums, t, m, memo)
		if ok {
			return ok
		}
	}
	return false
}

func canPartition(nums []int) bool {
	mask := 1<<len(nums) - 1
	memo := make(map[int]int)
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	return dfs(nums, sum/2, mask, memo)
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
	res := canPartition(nums)
	fmt.Println(res)
}
