package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type coordinate struct {
	x, y int
}

func (c coordinate) isZero() bool {
	if c.x == 0 && c.y == 0 {
		return true
	}
	return false
}

func getMoves(from coordinate, visited map[coordinate]struct{}) []coordinate {
	deltaX := []int{1, 2, 2, 1, -1, -2, -2, -1}
	deltaY := []int{-2, -1, 1, 2, 2, 1, -1, -2}
	// board is endless, so no need to check coordinates
	var next []coordinate
	for n := range deltaX {
		newC := coordinate{from.x + deltaX[n], from.y + deltaY[n]}
		if _, ok := visited[newC]; ok {
			continue
		}
		next = append(next, newC)
	}
	return next
}

func getKnightShortestPath(x int, y int) int {
	// WRITE YOUR BRILLIANT CODE HERE
	if x == 0 && y == 0 {
		return 0
	}
	var queue []coordinate
	queue = append(queue, coordinate{x, y})
	visited := make(map[coordinate]struct{})
	distance := 0
	for len(queue) > 0 {
		size := len(queue)
		distance++
		for i := 0; i < size; i++ {
			c := queue[0]
			queue = queue[1:]
			visited[c] = struct{}{}
			for _, m := range getMoves(c, visited) {
				if m.isZero() {
					return distance
				}
				queue = append(queue, m)
			}
		}
	}
	return distance
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	y, _ := strconv.Atoi(scanner.Text())
	res := getKnightShortestPath(x, y)
	fmt.Println(res)
}
