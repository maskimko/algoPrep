package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x, y int
}

func getNeighbors(grid [][]int, c coordinate) []coordinate {
	deltaX := []int{0, 1, 0, -1}
	deltaY := []int{-1, 0, 1, 0}
	var neighbors []coordinate
	for n := range deltaX {
		if c.x+deltaX[n] >= 0 && c.x+deltaX[n] < len(grid) {
			if c.y+deltaY[n] >= 0 && c.y+deltaY[n] < len(grid[c.x]) {
				if grid[c.x+deltaX[n]][c.y+deltaY[n]] > 0 {
					neighbors = append(neighbors, coordinate{c.x + deltaX[n], c.y + deltaY[n]})
				}
			}
		}
	}
	return neighbors
}

func visit(visited [][]bool, c coordinate) {
	visited[c.x][c.y] = true
}

func isVisited(visited [][]bool, c coordinate) bool {
	if visited[c.x] == nil {
		return false
	}
	return visited[c.x][c.y]
}

func bfs(queue []coordinate, grid [][]int, visited [][]bool) {
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			c := queue[0]
			queue = queue[1:]
			if !isVisited(visited, c) {
				visit(visited, c)
				for _, n := range getNeighbors(grid, c) {
					queue = append(queue, n)
				}
			}
		}
	}
}

func makeVisited(grid [][]int) [][]bool {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}
	return visited
}

func countNumberOfIslands(grid [][]int) int {
	// WRITE YOUR BRILLIANT CODE HERE
	visited := makeVisited(grid)
	var queue []coordinate
	islands := 0
	for j := range grid {
		for k := range grid[j] {
			c := coordinate{j, k}
			if !isVisited(visited, c) {
				if grid[j][k] > 0 {
					queue = append(queue, c)
					islands++
					bfs(queue, grid, visited)
				}
			}
		}
	}
	return islands
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
	gridLength, _ := strconv.Atoi(scanner.Text())
	grid := [][]int{}
	for i := 0; i < gridLength; i++ {
		scanner.Scan()
		grid = append(grid, arrayAtoi(splitWords(scanner.Text())))
	}
	res := countNumberOfIslands(grid)
	fmt.Println(res)
}
