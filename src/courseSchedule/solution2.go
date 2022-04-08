package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func makeGraph2(n int, prerequisites [][]int) map[int][]int {
	graph := make(map[int][]int)
	for i := 0; i < n; i++ {
		graph[i] = []int{}
	}
	for _, p := range prerequisites {
		for k := 1; k < len(p); k++ {
			graph[p[0]] = append(graph[p[0]], p[k])
		}
	}
	return graph
}

func dfs(root int, graph map[int][]int, visited map[int]struct{}, visiting map[int]struct{}) {
	if _, ok := visited[root]; ok {
		return
	}
	for _, e := range graph[root] {
		if _, ok := visiting[e]; ok {
			return // loop detected
		} else {
			visiting[e] = struct{}{}
			dfs(e, graph, visited, visiting)
		}
	}
	visited[root] = struct{}{}
	delete(visiting, root)
}

func isValidCourseSchedule2(n int, prerequisites [][]int) bool {
	graph := makeGraph2(n, prerequisites)
	visited := make(map[int]struct{})
	visiting := make(map[int]struct{})
	for i := 0; i < len(graph); i++ {
		if _, ok := visited[i]; !ok {
			visiting[i] = struct{}{}
			dfs(i, graph, visited, visiting)
		}
	}
	return len(visiting) == 0 && len(visited) == n
}

func main2() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	prerequisitesLength, _ := strconv.Atoi(scanner.Text())
	prerequisites := [][]int{}
	for i := 0; i < prerequisitesLength; i++ {
		scanner.Scan()
		prerequisites = append(prerequisites, arrayAtoi(splitWords(scanner.Text())))
	}
	res := isValidCourseSchedule2(n, prerequisites)
	fmt.Println(res)
}
