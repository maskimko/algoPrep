// [[a c l x],
//  [b l o x],
//  [x e h x],
//  [x x x x]]

// 2-dimentional array

// dictionary.is_word(str)
// 0 = no
// -1 = it's part of a word
// 1 = it's a complete word

package main

import (
	dictionary "algoMonster/src/interview/workato/dictionary"
	"fmt"
	"strings"
)

type coordinate struct {
	x int
	y int
}

func (c coordinate) getLeft() coordinate {
	return coordinate{c.x - 1, c.y}
}
func (c coordinate) getRight() coordinate {
	return coordinate{c.x + 1, c.y}
}
func (c coordinate) getUp() coordinate {
	return coordinate{c.x, c.y - 1}
}
func (c coordinate) getDown() coordinate {
	return coordinate{c.x, c.y + 1}
}

func validCoordinate(c coordinate, size int) bool {
	if c.y < 0 || c.y >= size {
		return false
	}
	if c.x < 0 || c.x >= size {
		return false
	}
	return true
}

func (c coordinate) equals(other *coordinate) bool {
	if other == nil {
		return false
	}
	if c.x == other.x && c.y == other.y {
		return true
	}
	return false
}

func neighbors(c, prev *coordinate, allowDiagonals bool, used [][]bool, data [][]rune) ([]coordinate, error) {
	// return up down left right
	size := len(data)
	if size == 0 {
		return nil, nil
	}
	if c == nil {
		return nil, fmt.Errorf("coordinate cannot be null")
	}
	if !validCoordinate(*c, size) {
		return nil, fmt.Errorf("coordinate [%d,%d] must be within the size of"+
			" the data grid (between 0 and %d)", c.x, c.y, len(data))
	}
	var coordinates []coordinate
	//left
	left := c.getLeft()
	if validCoordinate(left, size) && !isUsed(left, used) && !left.equals(prev) {
		coordinates = append(coordinates, left)
	}
	//right
	right := c.getRight()
	if validCoordinate(right, size) && !isUsed(right, used) && !right.equals(prev) {
		coordinates = append(coordinates, right)
	}
	//up
	up := c.getUp()
	if validCoordinate(up, size) && !isUsed(up, used) && !up.equals(prev) {
		coordinates = append(coordinates, up)
	}
	//down
	down := c.getDown()
	if validCoordinate(down, size) && !isUsed(down, used) && !down.equals(prev) {
		coordinates = append(coordinates, down)
	}
	//allow diagonals
	if allowDiagonals {
		//left up
		leftUp := left.getUp()
		if validCoordinate(leftUp, size) && !isUsed(leftUp, used) && !leftUp.equals(prev) {
			coordinates = append(coordinates, leftUp)
		}
		//right up
		rightUp := right.getUp()
		if validCoordinate(rightUp, size) && !isUsed(rightUp, used) && !rightUp.equals(prev) {
			coordinates = append(coordinates, rightUp)
		}
		//left down
		leftDown := left.getDown()
		if validCoordinate(leftDown, size) && !isUsed(leftDown, used) && !leftDown.equals(prev) {
			coordinates = append(coordinates, leftDown)
		}
		//right down
		rightDown := right.getDown()
		if validCoordinate(rightDown, size) && !isUsed(rightDown, used) && !rightDown.equals(prev) {
			coordinates = append(coordinates, rightDown)
		}
	}
	return coordinates, nil
}

func isUsed(c coordinate, used [][]bool) bool {
	if used == nil {
		return false
	}
	if used[c.x] == nil {
		used[c.x] = make([]bool, len(used))
	}
	return used[c.x][c.y]
}

func FindWord(data [][]rune) (string, error) {

	//populate rows

	for i := 0; i < len(data); i++ {
		for k := 0; k < len(data); k++ {
			used := make([][]bool, len(data))
			path := make([]coordinate, 0)
			word, err := findWord(&coordinate{i, k}, data, used, &path)
			if err != nil {
				return "", err
			}
			if word != "" {
				return word, nil
			}
		}
	}
	return "", nil
}

func isWord(c coordinate, data [][]rune, path *[]coordinate) int8 {
	if path == nil || len(*path) == 0 {
		return dictionary.IsWord(string(data[c.x][c.y]))
	}
	return dictionary.IsWord(coordinatesToString(append(*path, c), data))
}

func coordinatesToString(c []coordinate, data [][]rune) string {
	if len(c) == 0 {
		return ""
	}
	var chars []string
	for i := range c {
		chars = append(chars, string(data[c[i].x][c[i].y]))
	}
	return strings.Join(chars, "")
}

func findWord(c *coordinate, data [][]rune, used [][]bool, path *[]coordinate) (string, error) {

	wordResult := isWord(*c, data, path)
	if wordResult != 0 {
		if used[c.x] == nil {
			used[c.x] = make([]bool, len(data))
		}
		used[c.x][c.y] = true
		*path = append(*path, *c)
		if wordResult > 0 {
			return coordinatesToString(*path, data), nil
		}
		var prev *coordinate
		if path != nil && len(*path) > 0 {
			prev = &((*path)[len(*path)-1])
		}
		nbrs, err := neighbors(c, prev, true, used, data)
		if err != nil {
			return "", err
		}
		for _, n := range nbrs {
			w, err := findWord(&n, data, used, path)
			if err != nil {
				return "", err
			}
			if w != "" {
				return w, nil
			}
		}
	}
	return "", nil
}
