package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func makeGraph(n int, prerequisites [][]int) map[int][]int {
	graph := make(map[int][]int)
	for i := 0; i < n; i++ {
		graph[i] = []int{}
	}
	for i := 0; i < len(prerequisites); i++ {
		for k := 1; k < len(prerequisites[i]); k++ {
			graph[prerequisites[i][k]] = append(graph[prerequisites[i][k]], prerequisites[i][0])
		}
	}
	return graph
}

func countParents(n int, prerequisites [][]int) []int {
	counts := make([]int, n, n)
	for i := 0; i < n; i++ {
		counts[i] = 0
	}
	for i := 0; i < len(prerequisites); i++ {
		for k := 1; k < len(prerequisites[i]); k++ {
			counts[prerequisites[i][0]]++
		}
	}
	return counts
}

func isValidCourseSchedule(n int, prerequisites [][]int) bool {
	// WRITE YOUR BRILLIANT CODE HERE
	queue := make([]int, 0)
	graph := makeGraph(n, prerequisites)
	counts := countParents(n, prerequisites)
	for k := range counts {
		if counts[k] == 0 {
			queue = append(queue, k)
		}
	}
	visited := 0 // visited
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		for j := 0; j < len(graph[c]); j++ {
			counts[graph[c][j]]--
			if counts[graph[c][j]] == 0 {
				queue = append(queue, graph[c][j])
			}
		}
		visited++
	}
	return visited == n
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
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	prerequisitesLength, _ := strconv.Atoi(scanner.Text())
	prerequisites := [][]int{}
	for i := 0; i < prerequisitesLength; i++ {
		scanner.Scan()
		prerequisites = append(prerequisites, arrayAtoi(splitWords(scanner.Text())))
	}
	res := isValidCourseSchedule(n, prerequisites)
	fmt.Println(res)
}
