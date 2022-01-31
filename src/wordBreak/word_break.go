package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func wordBreak(s string, words []string) bool {
	cache := make(map[string]bool)
	return wordBreakCache(s, words, cache)
}

func wordBreakCache(s string, words []string, cache map[string]bool) bool {
	if len(s) == 0 || len(words) == 0 {
		return false
	}
	for i := range words {
		if len(s) < len(words[i]) {
			continue
		}
		if len(s) == len(words[i]) {
			if s == words[i] {
				cache[s] = true
				return true
			}
		}
		if res, ok := cache[s]; ok {
			return res
		}
		suffix := s[len(words[i]):]
		result := wordBreakCache(suffix, words, cache)
		if result {
			cache[s] = true
			return true
		}
	}
	cache[s] = false
	return false
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
	s := scanner.Text()
	scanner.Scan()
	words := splitWords(scanner.Text())
	res := wordBreak(s, words)
	fmt.Println(res)
}
