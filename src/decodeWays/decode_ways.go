package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func decodeWays(digits string) int {
	memo := make([]int, len(digits), len(digits))
	return decodeWaysMemo(0, digits, memo)
}

func decodeWaysMemo(i int, digits string, memo []int) int {
	if i == len(digits) {
		return 0
	}
	if len(memo) > i && memo[i] != 0 {
		return memo[i]
	}
	if i == len(digits)-1 {
		memo[i] = 1
		return 1
	}

	couple, err := strconv.Atoi(digits[i : i+2])
	if err != nil {
		return 0
	}
	if couple < 0 {
		return 0
	}

	ways := 0
	if couple < 26 && couple > 10 {
		if i == len(digits)-2 {
			memo[i] = 2
			return 2
		}
		ways += decodeWaysMemo(i+2, digits, memo)
	}
	if i == len(digits)-2 {
		memo[i] = 1
		return 1
	}
	ways += decodeWaysMemo(i+1, digits, memo)
	memo[i] = ways
	return ways
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	digits := scanner.Text()
	res := decodeWays(digits)
	fmt.Println(res)
}
