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

func initQueue(dMap [][]int) []coordinate {
	var queue []coordinate
	for j := range dMap {
		for k := range dMap[j] {
			if dMap[j][k] == 0 {
				queue = append(queue, coordinate{j, k})
			}
		}
	}
	return queue
}

func getNeighbors(dMap [][]int, c coordinate, distance int) []coordinate {
	deltaX := []int{0, 1, 0, -1}
	deltaY := []int{-1, 0, 1, 0}
	var neighbors []coordinate
	for n := range deltaX {
		xCoord := c.x + deltaX[n]
		if xCoord >= 0 && xCoord < len(dMap) {
			yCoord := c.y + deltaY[n]
			if yCoord >= 0 && yCoord < len(dMap[xCoord]) {
				if dMap[xCoord][yCoord] < 0 {
					continue
				}
				if dMap[xCoord][yCoord] > distance {
					neighbors = append(neighbors, coordinate{xCoord, yCoord})
				}
			}
		}
	}
	return neighbors
}

func mapGateDistances(dungeonMap [][]int) [][]int {
	// WRITE YOUR BRILLIANT CODE HERE
	queue := initQueue(dungeonMap)
	distance := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			c := queue[0]
			queue = queue[1:]
			if distance < dungeonMap[c.x][c.y] {
				dungeonMap[c.x][c.y] = distance
			}
			for _, n := range getNeighbors(dungeonMap, c, distance) {
				queue = append(queue, n)
			}
		}
		distance++
	}
	return dungeonMap
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

func arrayItoa(arr []int) []string {
	res := []string{}
	for _, v := range arr {
		res = append(res, strconv.Itoa(v))
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dungeonMapLength, _ := strconv.Atoi(scanner.Text())
	dungeonMap := [][]int{}
	for i := 0; i < dungeonMapLength; i++ {
		scanner.Scan()
		dungeonMap = append(dungeonMap, arrayAtoi(splitWords(scanner.Text())))
	}
	res := mapGateDistances(dungeonMap)
	for _, row := range res {
		fmt.Println(strings.Join(arrayItoa(row), " "))
	}
}
