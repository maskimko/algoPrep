package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func removeDuplicates(arr []int) int {
	var j = 0
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[j] {
			j++
			arr[j] = arr[i]
		}
	}
	return j + 1
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
	res := removeDuplicates(arr)
	fmt.Println(strings.Join(arrayItoa(arr[:res]), " "))
}
