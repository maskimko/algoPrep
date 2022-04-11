package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isAlphaNumeric(r rune) bool {
	if (r >= '0' && r <= '9') || (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
		return true
	}
	return false
}

func toLower(c rune) rune {
	return []rune(strings.ToLower(string(c)))[0]
}

func isPalindrome(s string) bool {
	chars := []rune(s)
	k := len(chars) - 1
	for i := 0; i < k; {
		if !isAlphaNumeric(chars[k]) {
			k--
			continue
		}
		if !isAlphaNumeric(chars[i]) {
			i++
			continue
		}
		if toLower(chars[i]) != toLower(chars[k]) {
			return false
		}
		k--
		i++
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	res := isPalindrome(s)
	fmt.Println(res)
}
