package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func shortestPath(graph [][]int, from, to int) int {
	// WRITE YOUR BRILLIANT CODE HERE
	visited := make(map[int]struct{})
	var queue []int
	distance := 0
	queue = append(queue, from)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			item := queue[0]
			queue = queue[1:]
			if _, ok := visited[item]; ok {
				continue
			}
			if item == to {
				return distance
			}
			neighbors := graph[item]
			for n := range neighbors {
				queue = append(queue, neighbors[n])
			}
		}
		distance++
	}
	return -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ok := scanner.Scan()
	if !ok {
		_, _ = fmt.Fprint(os.Stderr, "nothing to scan")
		return
	}
	graphLen, err := strconv.Atoi(scanner.Text())
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err.Error())
		return
	}
	graph := make([][]int, graphLen)
	for i := 0; i < graphLen; i++ {
		scanner.Scan()
		chunks := strings.Split(scanner.Text(), " ")
		for k := range chunks {
			neighbor, err := strconv.Atoi(chunks[k])
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "failed ot parse %v", chunks[k])
				return
			}
			graph[i] = append(graph[i], neighbor)
		}
	}
	scanner.Scan()
	from, err := strconv.Atoi(scanner.Text())
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to read from %s", err.Error())
		return
	}
	scanner.Scan()
	to, err := strconv.Atoi(scanner.Text())
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to read from %s", err.Error())
		return
	}
	res := shortestPath(graph, from, to)
	fmt.Println(res)
}
