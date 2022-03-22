package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func makeGraph(seqs [][]int) map[int][]int {
	graph := make(map[int][]int)
	for _, s := range seqs {
		for i := range s {
			if _, ok := graph[s[i]]; ok {
				continue
			}
			graph[s[i]] = []int{}
		}
	}
	for _, s := range seqs {
		for i := 0; i < len(s)-1; i++ {
			graph[s[i]] = append(graph[s[i]], s[i+1])
		}
	}
	return graph
}

func countParents(graph map[int][]int) map[int]int {
	counts := make(map[int]int)
	for k := range graph {
		counts[k] = 0
	}
	for _, v := range graph {
		for i := range v {
			counts[v[i]] = counts[v[i]] + 1
		}
	}
	return counts
}

func sequenceReconstruction(original []int, seqs [][]int) bool {
	// WRITE YOUR BRILLIANT CODE HERE
	var queue []int
	graph := makeGraph(seqs)
	parCount := countParents(graph)
	var result []int
	for k := range parCount {
		if parCount[k] == 0 {
			queue = append(queue, k)
		}
	}
	for len(queue) > 0 {
		if len(queue) > 1 {
			return false
		}
		c := queue[0]
		queue = queue[1:]
		result = append(result, c)
		for i := range graph[c] {
			parCount[graph[c][i]] = parCount[graph[c][i]] - 1
			if parCount[graph[c][i]] == 0 {
				queue = append(queue, graph[c][i])
			}
		}
	}
	if len(result) != len(original) {
		return false
	}
	for i := range result {
		if result[i] != original[i] {
			return false
		}
	}
	return true
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
	original := arrayAtoi(splitWords(scanner.Text()))
	scanner.Scan()
	seqsLength, _ := strconv.Atoi(scanner.Text())
	seqs := [][]int{}
	for i := 0; i < seqsLength; i++ {
		scanner.Scan()
		seqs = append(seqs, arrayAtoi(splitWords(scanner.Text())))
	}
	res := sequenceReconstruction(original, seqs)
	fmt.Println(res)
}
