package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func diff(one, other string) int {
	oneRunes := []rune(one)
	otherRunes := []rune(other)
	if len(one) != len(other) {
		return -1
	}
	counter := 0
	for n := range oneRunes {
		if oneRunes[n] != otherRunes[n] {
			counter++
		}
	}
	return counter
}

func wordLadder(begin string, end string, wordList []string) int {
	// WRITE YOUR BRILLIANT CODE HERE
	visited := make(map[string]bool)
	var queue []string
	queue = append(queue, begin)
	distance := 0
	for len(queue) > 0 {
		size := len(queue)
		distance++
		for i := 0; i < size; i++ {
			c := queue[0]
			queue = queue[1:]
			if _, ok := visited[c]; ok {
				continue
			}
			visited[c] = true
			for n := range wordList {
				if _, ok := visited[wordList[n]]; ok {
					continue
				}

				if diff(c, wordList[n]) == 1 {
					if wordList[n] == end {
						return distance
					}
					queue = append(queue, wordList[n])
				}
			}
		}
	}
	return distance
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
	begin := scanner.Text()
	scanner.Scan()
	end := scanner.Text()
	scanner.Scan()
	wordList := splitWords(scanner.Text())
	res := wordLadder(begin, end, wordList)
	fmt.Println(res)
}
