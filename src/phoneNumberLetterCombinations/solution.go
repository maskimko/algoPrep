package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var lettersMap []string = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinationsOfPhoneNumber(digits string) []string {
	result := make([]string, 0)
	combinations(digits, &[]rune{}, &result)
	return result
}

func combinations(input string, path *[]rune, result *[]string) {
	if len(input) == 0 {
		*result = append(*result, string(*path))
		return
	}

	number := input[0] - '0'
	letters := lettersMap[number]
	for j := range letters {
		*path = append(*path, rune(letters[j]))

		if len(input) > 0 {
			combinations(input[1:], path, result)
		}
		*path = (*path)[:len(*path)-1]
	}

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	digits := scanner.Text()
	res := letterCombinationsOfPhoneNumber(digits)
	fmt.Println(strings.Join(res, " "))
}
