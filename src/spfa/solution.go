package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type edge struct {
	to     int
	weight int
}

func shortestPath(graph map[int][]edge, root int, target int) int {
	// WRITE YOUR BRILLIANT CODE HERE

	return 0
}

func shortestPathStrInput(input string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	ok := scanner.Scan()
	if !ok {
		return -1, scanner.Err()
	}
	graphLen, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return -1, err
	}
	graph := make(map[int][]edge)
	for i := 0; i < graphLen; i++ {
		ok = scanner.Scan()
		if !ok {
			return -1, scanner.Err()
		}
		graphRecord := scanner.Text()
		edgePairs := strings.Split(graphRecord, " ")
		if len(edgePairs)%2 != 0 {
			return -1, fmt.Errorf("wrong input format of edge pairs of node %d", i)
		}
		var edges []edge
		for k := 0; k < len(edgePairs)-1; k += 2 {
			to, err := strconv.Atoi(edgePairs[k])
			if err != nil {
				return -1, fmt.Errorf("cannot parse egde %d connection of node %d", k, i)
			}
			weight, err := strconv.Atoi(edgePairs[k+1])
			if err != nil {
				return -1, fmt.Errorf("cannot parse weight of %d connection of node %d", k, i)
			}
			edges = append(edges, edge{to, weight})
		}
		graph[i] = edges
	}
	ok = scanner.Scan()
	if !ok {
		return -1, scanner.Err()
	}
	root, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return -1, fmt.Errorf("failed to parse root")
	}
	ok = scanner.Scan()
	if !ok {
		return -1, scanner.Err()
	}
	target, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return -1, fmt.Errorf("failed to parse root")
	}
	return shortestPath(graph, root, target), nil
}

func main() {
	all, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to read standard input, %s", err.Error())
	}

	res, err := shortestPathStrInput(string(all))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed toparse input, %s", err.Error())
	}
	fmt.Println(res)
}
