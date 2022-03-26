package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countParents(graph map[rune][]rune) map[rune]int {
	counts := make(map[rune]int)
	for r := range graph {
		counts[r] = 0
	}
	for r, deps := range graph {
		for i := 0; i < len(deps); i++ {
			if deps[i] != r {
				counts[deps[i]] = counts[deps[i]] + 1
			}
		}
	}
	return counts
}

func makeGraph(words []string) map[rune][]rune {
	graph := make(map[rune][]rune)
	for _, w := range words {
		runes := []rune(w)
		for i := 1; i < len(runes); i++ {
			if runes[i-1] != runes[i] {
				graph[runes[i-1]] = append(graph[runes[i-1]], runes[i])
			}
			if _, ok := graph[runes[i]]; !ok {
				graph[runes[i]] = nil
			}
		}
	}
	return graph
}

func alienOrder(words []string) string {
	graph := makeGraph(words)
	parCount := countParents(graph)
	visited := make(map[rune]bool)
	queue := make([]rune, 0)
	var result []rune
	for r := range parCount {
		if parCount[r] == 0 {
			queue = append(queue, r)
		}
	}
	for len(queue) > 0 {
		//if len(queue)!=1{
		//	// undetermined order
		//	return ""
		//}
		c := queue[0]
		queue = queue[1:]
		if visited[c] {
			continue
		}
		visited[c] = true
		result = append(result, c)
		for i := 0; i < len(graph[c]); i++ {
			parCount[graph[c][i]] -= 1
			if parCount[graph[c][i]] == 0 {
				queue = append(queue, graph[c][i])
			}
		}
	}
	return string(result)
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
	words := splitWords(scanner.Text())
	res := alienOrder(words)
	fmt.Println(res)
}
