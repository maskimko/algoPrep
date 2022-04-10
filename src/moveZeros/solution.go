package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func moveZeros(nums []int) {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[k] != 0 {
			k++
			continue
		}
		if nums[i] != 0 {
			nums[k], nums[i] = nums[i], nums[k]
			k++
		}
	}
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
	nums := arrayAtoi(splitWords(scanner.Text()))
	moveZeros(nums)
	fmt.Println(strings.Join(arrayItoa(nums), " "))
}
