package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type distance struct {
	node  int
	value int
	from  int
}

type pq []distance

func (p pq) Len() int {
	return len(p)
}

func (p pq) Less(i, j int) bool {
	return p[i].value < p[j].value
}

func (p pq) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *pq) Push(x interface{}) {
	d := x.(distance)
	*p = append(*p, d)
}

func (p *pq) Pop() interface{} {
	d := (*p)[len(*p)-1]
	*p = (*p)[0 : len(*p)-1]
	return d
}

type edge struct {
	to     int
	weight int
}

func shortestDejkstraPath(graph [][]edge, root, target int) (int, []int) {
	distances := make([]distance, len(graph))
	intMax := 1<<32 - 1
	for i := 0; i < len(distances); i++ {
		distances[i] = distance{i, intMax, -1}
	}
	distances[root].value = 0
	queue := make(pq, 0)
	heap.Push(&queue, distances[root])
	var targetDist distance
	for len(queue) > 0 {
		c := heap.Pop(&queue).(distance)
		if c.node == target {
			targetDist = c
			break
		}
		for k := 0; k < len(graph[c.node]); k++ {
			n := graph[c.node][k]
			if distances[n.to].value <= n.weight+c.value {
				continue
			}
			distances[n.to].value = n.weight + c.value
			distances[n.to].from = c.node
			heap.Push(&queue, distances[n.to])
		}
	}
	pathLen := targetDist.value
	path := make([]int, 0)
	for d := targetDist; d.from != root; d = distances[d.from] {
		path = append([]int{d.node}, path...)
	}
	path = append([]int{root}, path...)
	return pathLen, path
}

func shortestPath(graph [][]edge, root, target int) int {
	d, _ := shortestDejkstraPath(graph, root, target)
	return d
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
	graph := make([][]edge, graphLen)
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
