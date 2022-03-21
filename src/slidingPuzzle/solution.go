package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func detectZero(pos [][]int) (x int, y int) {
	for x = 0; x < len(pos); x++ {
		for y = 0; y < len(pos[x]); y++ {
			if pos[x][y] == 0 {
				return x, y
			}
		}
	}
	return -1, -1
}

func swap(x, y, dx, dy int, pos [][]int) {
	tmp := pos[x+dx][y+dy]
	pos[x+dx][y+dy] = pos[x][y]
	pos[x][y] = tmp
}

func clone(pos [][]int) [][]int {
	dup := make([][]int, len(pos))
	for n := range pos {
		dup[n] = make([]int, len(pos[n]))
		copy(dup[n], pos[n])
	}
	return dup
}

func key(pos [][]int) int {
	k := 0
	for i := range pos {
		for j := range pos[i] {
			k += k*10 + pos[i][j]
		}
	}
	return k
}

func neighbors(pos [][]int) [][][]int {
	x, y := detectZero(pos)
	var result [][][]int
	if x > 0 {
		dup := clone(pos)
		swap(x, y, -1, 0, dup)
		result = append(result, dup)
	}
	if x < len(pos)-1 {
		dup := clone(pos)
		swap(x, y, 1, 0, dup)
		result = append(result, dup)
	}
	if y > 0 {
		dup := clone(pos)
		swap(x, y, 0, -1, dup)
		result = append(result, dup)
	}
	if y < len(pos[x])-1 {
		dup := clone(pos)
		swap(x, y, 0, 1, dup)
		result = append(result, dup)
	}
	return result
}

func isGoal(pos [][]int) bool {
	k := 0
	for i := range pos {
		for j := range pos[i] {
			k++
			if i == len(pos)-1 && j == len(pos[i])-1 && pos[i][j] == 0 {
				return true
			}
			if k != pos[i][j] {
				return false
			}
		}
	}
	return true
}

func numSteps(initPos [][]int) int {
	// WRITE YOUR BRILLIANT CODE HERE
	var queue [][][]int
	visited := make(map[int]struct{})
	queue = append(queue, initPos)
	distance := 0
	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			c := queue[0]
			queue = queue[1:]
			if _, ok := visited[key(c)]; ok {
				continue
			}
			visited[key(c)] = struct{}{}
			if isGoal(c) {
				return distance
			}
			for _, n := range neighbors(c) {
				queue = append(queue, n)
			}
		}
		distance++
		if distance > len(initPos)*len(initPos[0])*5 {
			return -1
		}
	}
	return -1
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
	initPosLength, _ := strconv.Atoi(scanner.Text())
	initPos := [][]int{}
	for i := 0; i < initPosLength; i++ {
		scanner.Scan()
		initPos = append(initPos, arrayAtoi(splitWords(scanner.Text())))
	}
	res := numSteps(initPos)
	fmt.Println(res)
}
