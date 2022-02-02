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
	"dictionary"
)

type coordinate struct {
	x int
	y int
}

func neighbors(c, prev *coordinate, data [][]string) []coordinates {
	// return up down left right
	return nil
}

func isVisited(c coordinate, visited [][]bool) bool {
	return visited[c.x][c.y]
}

func FindWord(data [][]string) {
	visited := make([][]bool,len(data))
	//populate rows
	path := make([]string,0)
	var c coordinate = coordinate{0,0} //random is better
	findWord(data, visited, path)
}


func findWord(c coordinate, data [][]string, visited [][]bool, path *[]coordinate) string {
	nbrs := neighbors(c, nil, data)
	if dictionary.isWord(data[c.x][c.y]) {
		return
	}
	///append to a path and check neighbors
	{
		//update visited USED
		*path := append(*path, data[x][y])
		for _, n := range neighbors {
			if !isVisited(n, visited) {
				findWord(n, data, visited, path)
			}
		}

	}
}

}

