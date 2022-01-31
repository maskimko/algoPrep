package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func newspapersSplit(newspapers []int, coworkers int) int {
	// WRITE YOUR BRILLIANT CODE HERE
	if coworkers <= 0 {
		return -1
	}
	if len(newspapers) == 0 {
		return 0
	}
	if len(newspapers) > 100000 || coworkers > 100000 {
		return -1
	}
	low := 1
	high := 1<<31 - 2 //just to prevent overflow
	var mid int = 0
	for low <= high {
		mid = low + (high-low)/2
		result := check(newspapers, coworkers, mid)
		if result {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return high + 1
}

func check(newspapers []int, coworkers int, daysLimit int) bool {
	sum := 0
	split := 0
	for i := 0; i < len(newspapers); i++ {
		if newspapers[i] > daysLimit {
			return false
		}
		sum += newspapers[i]
		if sum > daysLimit {
			split++
			sum = newspapers[i]
		}
	}
	return split+1 <= coworkers
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
	newspapers := arrayAtoi(splitWords(scanner.Text()))
	scanner.Scan()
	coworkers, _ := strconv.Atoi(scanner.Text())
	res := newspapersSplit(newspapers, coworkers)
	fmt.Println(res)
}
