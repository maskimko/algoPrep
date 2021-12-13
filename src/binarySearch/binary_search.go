package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func binarySearch(arr []int, target int) int {
	// WRITE YOUR BRILLIANT CODE HERE

	if len(arr) == 0 {
		return -1
	}
	i := len(arr) / 2
	j := len(arr)
	for {
		if i >= j {
			break
		}
		if arr[i] == target {
			return i
		}
		if arr[i] > target {
			j = i
			i /= 2
		} else {
			i += (j - i) / 2
		}
	}

	return -1
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
	arr := arrayAtoi(splitWords(scanner.Text()))
	scanner.Scan()
	target, _ := strconv.Atoi(scanner.Text())
	res := binarySearch(arr, target)
	fmt.Println(res)
}
