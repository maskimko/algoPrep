package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type task struct {
	id           string
	parents      int
	deps         []*task
	requiredTime int
	visited      bool
	spentTime    int
}

func countParents(graph map[string]*task) {
	for _, t := range graph {
		t.parents = 0
	}
	for _, t := range graph {
		for i := range t.deps {
			t.deps[i].parents++
		}
	}
}

func buildGraph(tasks []string, times []int, requirements [][]string) map[string]*task {
	taskMap := make(map[string]*task)
	for i := range tasks {
		taskMap[tasks[i]] = &task{id: tasks[i], requiredTime: times[i]}
	}
	for _, r := range requirements {
		for i := 1; i < len(r); i++ {
			taskMap[r[i]].deps = append(taskMap[r[i]].deps, taskMap[r[i-1]])
		}
	}
	return taskMap
}

func taskScheduling2(tasks []string, times []int, requirements [][]string) int {
	graph := buildGraph(tasks, times, requirements)
	countParents(graph)
	queue := make([]string, 0)
	totalTime := 0
	for k, t := range graph {
		if t.parents == 0 {
			queue = append(queue, k)
			t.spentTime = t.requiredTime
		}
	}

	for len(queue) > 0 {
		size := len(queue)
		for n := 0; n < size; n++ {
			c := queue[0]
			queue = queue[1:]
			if graph[c].visited {
				continue
			}
			graph[c].visited = true
			if totalTime < graph[c].spentTime {
				totalTime = graph[c].spentTime
			}
			for _, j := range graph[c].deps {
				graph[j.id].parents--
				if graph[j.id].parents == 0 {
					if graph[j.id].spentTime < graph[c].spentTime+graph[j.id].requiredTime {
						graph[j.id].spentTime += graph[c].spentTime + graph[j.id].requiredTime
					}
					queue = append(queue, j.id)
				}
			}
		}
	}
	return totalTime
}

func splitWords(s string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, " ")
}

func arrayAtoi(arr []string) []int {
	res := []int{}
	for _, x := range arr {
		v, _ := strconv.Atoi(x)
		res = append(res, v)
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	tasks := splitWords(scanner.Text())
	scanner.Scan()
	times := arrayAtoi(splitWords(scanner.Text()))
	scanner.Scan()
	requirementsLength, _ := strconv.Atoi(scanner.Text())
	requirements := [][]string{}
	for i := 0; i < requirementsLength; i++ {
		scanner.Scan()
		requirements = append(requirements, splitWords(scanner.Text()))
	}
	res := taskScheduling2(tasks, times, requirements)
	fmt.Println(res)
}
