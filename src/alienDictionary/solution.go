package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strings"
)

type pQueue []rune

func (p pQueue) Len() int {
	return len(p)
}

func (p pQueue) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p pQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *pQueue) Push(x interface{}) {
	*p = append(*p, x.(rune))
}

func (p *pQueue) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

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
	if len(words) == 0 {
		return nil
	}
	for _, w := range words {
		for _, r := range []rune(w) {
			graph[r] = nil
		}
	}
	for n := 1; n < len(words); n++ {
		bigger := []rune(words[n-1])
		smaller := []rune(words[n])
		for i := 0; i < len(bigger) && i < len(smaller); i++ {
			if bigger[i] == smaller[i] {
				continue
			}
			graph[bigger[i]] = append(graph[bigger[i]], smaller[i])
			break
		}
	}
	return graph
}

func alienOrder(words []string) string {
	graph := makeGraph(words)
	parCount := countParents(graph)
	visited := make(map[rune]bool)
	pq := &pQueue{}
	heap.Init(pq)
	var result []rune
	for t := 0; pq.Len() == 0; t++ {
		for r := range parCount {
			if parCount[r] == t {
				heap.Push(pq, r)
			}
		}
	}
	for pq.Len() > 0 {
		c := heap.Pop(pq).(rune)
		if visited[c] {
			continue
		}
		visited[c] = true
		result = append(result, c)
		for i := 0; i < len(graph[c]); i++ {
			parCount[graph[c][i]] -= 1
			if parCount[graph[c][i]] == 0 {
				heap.Push(pq, graph[c][i])
			}
		}
	}
	if len(result) < len(graph) {
		return ""
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
