package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func twoSumSorted(arr []int, target int) []int {

	if len(arr) == 0 {
		return []int{}
	}
	k := len(arr) - 1
	for i := 0; i <= k; {
		sum := arr[i] + arr[k]
		if sum == target {
			return []int{i, k}
		}
		if sum > target {
			k--
		} else {
			i++
		}
	}
	return []int{}
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

func arrayItoa(arr []int) []string {
	res := []string{}
	for _, v := range arr {
		res = append(res, strconv.Itoa(v))
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	arr := arrayAtoi(splitWords(scanner.Text()))
	scanner.Scan()
	target, _ := strconv.Atoi(scanner.Text())
	res := twoSumSorted(arr, target)
	fmt.Println(strings.Join(arrayItoa(res), " "))
}
