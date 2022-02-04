package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func partition(s string) [][]string {
	var result [][]string
	var path []string
	walk(0, s, &path, &result)
	var reverse = make([][]string, len(result))
	for i := 0; i < len(result); i++ {
		reverse[i] = result[len(result)-1-i]
	}
	return reverse
}

func walk(i int, s string, path *[]string, result *[][]string) {
	if i >= len(s) {
		p := make([]string, len(*path), len(*path))
		copy(p, *path)
		*result = append(*result, p)
		return
	}
	for j := len(s) - 1; j > i; j-- {
		if s[j] == s[i] {
			if isPalindrome(s[i : j+1]) {
				pl := len(*path)
				*path = append(*path, s[i:j+1])
				walk(j+1, s, path, result)
				*path = (*path)[0:pl]
			}
		}
	}
	*path = append(*path, s[i:i+1])
	walk(i+1, s, path, result)
}

func backtrack(i, j int, s string, path *[]string, result *[][]string) {
	if i >= j {
		return
	}
	m := (j - i) / 2
	for w := 0; m-w > i; w++ {
		r := isPalindrome(s[m-w : m+1+w])
		if r {
			*path = append(*path, s[m-w:m+1+w])
			backtrack(i, m-w, s, path, result)
			backtrack(m+1+w, j, s, path, result)
		} else {
			break
		}
	}
	backtrack(i+1, j, s, path, result)
	backtrack(i, j-1, s, path, result)
}

func isPalindrome(p string) bool {
	if len(p) == 0 {
		return false
	}
	if len(p) == 1 {
		return true
	}
	for k := 0; k < len(p)/2; k++ {
		if p[k] != p[len(p)-k-1] {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	res := partition(s)
	for _, row := range res {
		fmt.Println(strings.Join(row, " "))
	}
}
