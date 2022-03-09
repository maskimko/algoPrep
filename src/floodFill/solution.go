package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x, y int
}

func getOrdinal(c coordinate, size int) int {
	return c.y*size + c.x
}

func getNeighbors(c coordinate, size int) []coordinate {
	deltaX := []int{0, 1, 0, -1}
	deltaY := []int{-1, 0, 1, 0}
	var neighbors []coordinate
	for i := 0; i < len(deltaX); i++ {
		nx := c.x + deltaX[i]
		ny := c.y + deltaY[i]
		if nx >= 0 && nx < size {
			if ny >= 0 && ny < size {
				neighbors = append(neighbors, coordinate{
					x: nx,
					y: ny,
				})
			}
		}
	}
	return neighbors
}

func getVal(image [][]int, c coordinate) int {
	return image[c.x][c.y]
}

func setVal(image [][]int, c coordinate, replacement int) {
	image[c.x][c.y] = replacement
}

func floodFill(r int, c int, replacement int, image [][]int) [][]int {
	// WRITE YOUR BRILLIANT CODE HERE
	val := getVal(image, coordinate{r, c})
	dim := len(image)
	var queue []coordinate
	queue = append(queue, coordinate{r, c})
	visited := make([]bool, dim*dim)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			crd := queue[0]
			queue = queue[1:]
			if visited[getOrdinal(crd, dim)] {
				continue
			}
			visited[getOrdinal(crd, dim)] = true

			if getVal(image, crd) == val {
				setVal(image, crd, replacement)
				neighbors := getNeighbors(crd, len(image))
				for n := range neighbors {
					if !visited[getOrdinal(neighbors[n], dim)] {
						queue = append(queue, neighbors[n])
					}
				}
			}
		}
	}
	return image
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

func processInput(reader io.Reader) (int, int, int, [][]int) {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	r, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	c, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	replacement, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	imageLength, _ := strconv.Atoi(scanner.Text())
	image := [][]int{}
	for i := 0; i < imageLength; i++ {
		scanner.Scan()
		image = append(image, arrayAtoi(splitWords(scanner.Text())))
	}
	return r, c, replacement, image
}

func main() {
	r, c, replacement, image := processInput(os.Stdin)
	res := floodFill(r, c, replacement, image)
	for _, row := range res {
		fmt.Println(strings.Join(arrayItoa(row), " "))
	}
}
