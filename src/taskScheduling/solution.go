package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func makeParentGraph(tasks []string, requirements [][]string) map[string][]string {
	graph := make(map[string][]string)
	for _, t := range tasks {
		graph[t] = []string{}
	}
	for _, r := range requirements {
		for i := 1; i < len(r); i++ {
			graph[r[0]] = append(graph[r[0]], r[i])
		}
	}
	return graph
}

func makeGraph(tasks []string, requirements [][]string) map[string][]string {
	graph := make(map[string][]string)
	for _, t := range tasks {
		graph[t] = []string{}
	}
	for _, r := range requirements {
		for i := 1; i < len(r); i++ {
			graph[r[i]] = append(graph[r[i]], r[0])
		}
	}
	return graph
}

func countParents(tasks []string, requirements [][]string) map[string]int {
	taskCount := make(map[string]int)
	for _, t := range tasks {
		taskCount[t] = 0
	}
	for _, r := range requirements {
		for i := 1; i < len(r); i++ {
			taskCount[r[i]] = taskCount[r[i]] + 1
		}
	}
	return taskCount
}

func taskScheduling(tasks []string, requirements [][]string) []string {
	// WRITE YOUR BRILLIANT CODE HERE
	var queue []string
	var result []string
	refCounts := countParents(tasks, requirements)
	graph := makeParentGraph(tasks, requirements)
	for k, v := range refCounts {
		if v == 0 {
			queue = append(queue, k)
		}
	}
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		result = append(result, c)
		for n := range graph[c] {
			refCounts[graph[c][n]] = refCounts[graph[c][n]] - 1
			if refCounts[graph[c][n]] == 0 {
				queue = append(queue, graph[c][n])
			}
		}
	}
	return result
}

func dfs(node string, graph map[string][]string, visited map[string]struct{}, result *[]string) {
	for n := range graph[node] {
		dfs(graph[node][n], graph, visited, result)
	}
	if _, ok := visited[node]; !ok {
		*result = append(*result, node)
		visited[node] = struct{}{}
	}
}

func taskSchedulingDFS(tasks []string, requirements [][]string) []string {
	graph := makeGraph(tasks, requirements)
	var result []string
	visited := make(map[string]struct{})
	for k := range graph {
		if _, ok := visited[k]; !ok {
			dfs(k, graph, visited, &result)
		}
	}
	return result
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
	tasks := splitWords(scanner.Text())
	scanner.Scan()
	requirementsLength, _ := strconv.Atoi(scanner.Text())
	requirements := [][]string{}
	for i := 0; i < requirementsLength; i++ {
		scanner.Scan()
		requirements = append(requirements, splitWords(scanner.Text()))
	}
	res := taskScheduling(tasks, requirements)
	fmt.Println(strings.Join(res, " "))
}
